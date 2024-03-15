package http

import (
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
	"server/bot/http_cmds"
	"server/db"
	"slices"
)

func ChannelsAdminAccept(writer http.ResponseWriter, request *http.Request) {
	parsedBody := parseBody(request, &writer)

	if parsedBody["channels_id"] == nil || parsedBody["chat_id"] == nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("There is no id field"))
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

	channelsByAdmin, err := db.New(conn).ChannelsByAdmin(db.Ctx, pgtype.Int8{
		Int64: int64(parsedBody["chat_id"].(float64)),
		Valid: true,
	})

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while check admins"))
		println(err.Error())
		return
	}

	if slices.IndexFunc(channelsByAdmin, func(admin db.ChannelsAdmin) bool {
		return admin.ChannelsID.Int32 == int32(parsedBody["channels_id"].(float64))
	}) == -1 {
		err = db.New(conn).InsertChannelAdmins(db.Ctx, db.InsertChannelAdminsParams{
			ChatID: pgtype.Int8{
				Int64: int64(parsedBody["chat_id"].(float64)),
				Valid: true,
			},
			ChannelsID: pgtype.Int4{
				Int32: int32(parsedBody["channels_id"].(float64)),
				Valid: true,
			},
		})

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while try to accept request"))
			println(err.Error())
			return
		}

		gl, err := db.New(conn).ChannelById(db.Ctx, int64(parsedBody["channels_id"].(float64)))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error while try to get group list"))
			println(err.Error())
			return
		}

		http_cmds.SendAcceptOfRequestOfAdminChannels(int64(parsedBody["chat_id"].(float64)), gl)
	}

	err = db.New(conn).DeleteListAdminsChannelRequestByChannelAndChatId(db.Ctx, db.DeleteListAdminsChannelRequestByChannelAndChatIdParams{
		ChatID: pgtype.Int8{
			Int64: int64(parsedBody["chat_id"].(float64)),
			Valid: true,
		},
		ChannelsID: pgtype.Int4{
			Int32: int32(parsedBody["channels_id"].(float64)),
			Valid: true,
		},
	})

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error while try to clear requests"))
		println(err.Error())
		return
	}
}
