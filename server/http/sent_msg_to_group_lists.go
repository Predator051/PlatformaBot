package http

import (
	"net/http"
	"server/bot/http_cmds"
	"server/db"
)

func SentMsgToChannelsRequest(writer http.ResponseWriter, request *http.Request) {
	parsedBody := parseBody(request, &writer)

	if parsedBody["channels_id"] == nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("There is no channels_id field"))
		return
	}

	if parsedBody["msg"] == nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("There is no msg field"))
		return
	}

	conn, err := db.NewConn()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while connect to db"))
		println(err.Error())
		return
	}

	defer conn.Close(db.Ctx)

	grl, err := db.New(conn).ChannelById(db.Ctx, int64(parsedBody["channels_id"].(float64)))

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while connect to db"))
		println(err.Error())
		return
	}

	err = http_cmds.SendMsgToChannel(conn, parsedBody["msg"].(string), grl)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while connect to db"))
		println(err.Error())
		return
	}
}
