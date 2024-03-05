package cmds

import (
	"fmt"
	"github.com/SakoDroid/telego/v2"
	"github.com/SakoDroid/telego/v2/objects"
	"github.com/jackc/pgx/v5/pgtype"
	"server/db"
	"strconv"
)

func RequestForAdminsPrivileges(bot *telego.Bot, update *objects.Update) {
	println("RequestForAdminsPrivileges start")

	conn, err := db.NewConn()

	if err != nil {
		bot.SendMessage(
			update.Message.Chat.Id,
			"Can't create db connection",
			"",
			update.Message.MessageId,
			false,
			false)
		return
	}

	defer conn.Close(db.Ctx)

	groupLists, err := db.New(conn).ListGroupList(db.Ctx)

	if err != nil {
		bot.SendMessage(
			update.Message.Chat.Id,
			"Can't get group news list",
			"",
			update.Message.MessageId,
			false,
			false)
	}

	kb := bot.CreateInlineKeyboard()

	for i, groupList := range groupLists {
		kb.AddCallbackButtonHandler(groupList.Name, strconv.FormatInt(groupList.ID, 10), i+1, func(u *objects.Update) {
			c, _ := db.NewConn()

			defer c.Close(db.Ctx)

			chat := u.CallbackQuery.Message.Chat
			groupListID, _ := strconv.ParseInt(u.CallbackQuery.Data, 10, 64)

			db.New(c).InsertListAdminsGroupListRequest(db.Ctx, db.InsertListAdminsGroupListRequestParams{
				Username: pgtype.Text{
					String: chat.Username,
					Valid:  true,
				},
				ChatID: pgtype.Int8{
					Int64: int64(chat.Id),
					Valid: true,
				},
				FirstName: pgtype.Text{
					String: chat.FirstName,
					Valid:  true,
				},
				SecondName: pgtype.Text{
					String: chat.LastName,
					Valid:  true,
				},
				GroupListID: pgtype.Int4{
					Int32: int32(groupListID),
					Valid: true,
				},
			})

			bot.SendMessage(
				chat.Id,
				"Your request was sent to our admins. Wait for answer from us for passing!",
				"",
				0,
				false,
				false)
		})
	}

	_, err = bot.AdvancedMode().ASendMessage(update.Message.Chat.Id,
		"Select group which you want to request",
		"",
		update.Message.MessageId,
		0,
		false,
		false,
		nil,
		false,
		false,
		kb)
	if err != nil {
		fmt.Println(err)
	}
}
