package http

import (
	"encoding/json"
	"net/http"
	"server/db"
)

func GroupListsSubscriptions(writer http.ResponseWriter, request *http.Request) {
	conn, err := db.NewConn()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while connect to db"))
		return
	}

	defer conn.Close(db.Ctx)

	groupList, err := db.New(conn).SubscriptionToGroupLists(db.Ctx)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while fetch group lists requests: " + err.Error()))
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
}
