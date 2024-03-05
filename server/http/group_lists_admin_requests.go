package http

import (
	"encoding/json"
	"net/http"
	"server/db"
	"slices"
)

func GroupListsAdminRequestsRequest(writer http.ResponseWriter, request *http.Request) {
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
			return list.ID == int64(listRequest.GroupListID.Int32)
		})

		if indx >= 0 {
			continue
		}

		grList, err := db.New(conn).GroupListById(db.Ctx, int64(listRequest.GroupListID.Int32))
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
}
