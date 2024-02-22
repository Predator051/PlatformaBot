package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"math/rand"
	"net/http"
	"server/db"
	"server/db/redis"
	"slices"
)

var AuthToken = "Bius2019!"

func parseBody(r *http.Request, w *http.ResponseWriter) map[string]any {
	parsedBody := map[string]any{}
	err := json.NewDecoder(r.Body).Decode(&parsedBody)
	if err != nil {
		http.Error(*w, err.Error(), http.StatusBadRequest)
		return map[string]any{}
	}

	return parsedBody
}

func randSession() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 20)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func NewServer(port int) {
	r := mux.NewRouter()

	r.HandleFunc("/auth", func(writer http.ResponseWriter, request *http.Request) {
		parsedBody := parseBody(request, &writer)

		if parsedBody["token"] != nil && parsedBody["token"] == AuthToken {
			session := randSession()

			err := redis.SetSession(session)

			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte(err.Error()))
				return
			}

			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(session))
		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("There is no token!"))
		}

		return
	}).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/group_lists", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := db.NewConn()

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while connect to db"))
			return
		}

		defer conn.Close(db.Ctx)

		groupList, err := db.New(conn).ListGroupList(db.Ctx)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while fetch group lists: " + err.Error()))
			return
		}

		responseData, err := json.Marshal(groupList)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while preparing response"))
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(responseData)
	}).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/new/group_lists", func(writer http.ResponseWriter, request *http.Request) {

		parsedBody := parseBody(request, &writer)

		if parsedBody["name"] == nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("There is no name field"))
			return
		}

		conn, err := db.NewConn()

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while connect to db"))
			return
		}

		defer conn.Close(db.Ctx)

		err = db.New(conn).InsertNewGroupList(db.Ctx, parsedBody["name"].(string))

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while fetch group lists: " + err.Error()))
			return
		}

		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/delete/group_lists", func(writer http.ResponseWriter, request *http.Request) {

		parsedBody := parseBody(request, &writer)

		if parsedBody["id"] == nil {
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte("There is no id field"))
			return
		}

		conn, err := db.NewConn()

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while connect to db"))
			return
		}

		defer conn.Close(db.Ctx)

		err = db.New(conn).DeleteGroupList(db.Ctx, int64(parsedBody["id"].(float64)))

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while fetch group lists: " + err.Error()))
			return
		}

		writer.WriteHeader(http.StatusOK)
	}).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/api/group_lists/admin/requests", func(writer http.ResponseWriter, request *http.Request) {
		conn, err := db.NewConn()

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while connect to db"))
			println(err.Error())
			return
		}

		defer conn.Close(db.Ctx)

		groupListRequests, err := db.New(conn).ListAdminsGroupListRequest(db.Ctx)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while fetch group lists requests: " + err.Error()))
			return
		}

		groupLists := []db.GroupList{}
		for _, listRequest := range groupListRequests {
			indx := slices.IndexFunc(groupLists, func(list db.GroupList) bool {
				return list.ID == listRequest.GroupListID.Int64
			})

			if indx >= 0 {
				continue
			}

			grList, err := db.New(conn).GroupListById(db.Ctx, listRequest.GroupListID.Int64)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte("Error while fetch group lists requests: " + err.Error()))
				return
			}

			groupLists = append(groupLists, grList)
		}

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while fetch group lists requests: " + err.Error()))
			return
		}

		responseData, err := json.Marshal(struct {
			Requests   []db.AdminsOfGroupListRequest `json:"requests"`
			GroupLists []db.GroupList                `json:"groupLists"`
		}{groupListRequests, groupLists})

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while preparing response"))
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Write(responseData)
	}).Methods(http.MethodPost, http.MethodOptions)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:8080", "*"},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type"},
	}).Handler(r)

	http.ListenAndServe(fmt.Sprintf("%s:%d", "127.0.0.1", port), handler)
}
