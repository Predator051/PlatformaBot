package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"math/rand"
	"net/http"
	"server/helper"
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

	r.HandleFunc("/auth", AuthRequest).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/channels", ChannelsRequest).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/new/channels", NewGroupListRequest).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/delete/channels", DeleteChannelsRequest).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/channels/admin/requests", ChannelsAdminRequestsRequest).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/channels/admin/accept", ChannelsAdminAccept).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/channels/subscriptions", ChannelsSubscriptions).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/api/channels/send/msg", SentMsgToChannelsRequest).Methods(http.MethodPost, http.MethodOptions)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:8080", "*"},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type"},
	}).Handler(r)

	log.Println(fmt.Sprintf("Started server on %s:%d", helper.SERVER_IP, port))
	http.ListenAndServe(fmt.Sprintf("%s:%d", helper.SERVER_IP, port), handler)
}
