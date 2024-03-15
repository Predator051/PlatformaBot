package http

import (
	"encoding/json"
	"net/http"
	"server/db"
	"slices"
)

func ChannelsAdminRequestsRequest(writer http.ResponseWriter, request *http.Request) {
	conn, err := db.NewConn()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while connect to db"))
		println(err.Error())
		return
	}

	defer conn.Close(db.Ctx)

	groupListRequests, err := db.New(conn).ListAdminsChannelRequest(db.Ctx)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while fetch group lists requests: " + err.Error()))
		return
	}

	channels := []db.Channel{}
	for _, listRequest := range groupListRequests {
		indx := slices.IndexFunc(channels, func(list db.Channel) bool {
			return list.ID == int64(listRequest.ChannelsID.Int32)
		})

		if indx >= 0 {
			continue
		}

		grList, err := db.New(conn).ChannelById(db.Ctx, int64(listRequest.ChannelsID.Int32))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while fetch group lists requests: " + err.Error()))
			return
		}

		channels = append(channels, grList)
	}

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while fetch group lists requests: " + err.Error()))
		return
	}

	responseData, err := json.Marshal(struct {
		Requests []db.AdminsOfChannelsRequest `json:"requests"`
		Channels []db.Channel                 `json:"channels"`
	}{groupListRequests, channels})

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while preparing response"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(responseData)
}
