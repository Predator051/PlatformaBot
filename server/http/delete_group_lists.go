package http

import (
	"net/http"
	"server/db"
)

func DeleteGroupListsRequest(writer http.ResponseWriter, request *http.Request) {

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
}
