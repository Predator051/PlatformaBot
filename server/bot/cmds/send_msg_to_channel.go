package cmds

import (
	"fmt"
	"github.com/SakoDroid/telego/v2"
	"github.com/SakoDroid/telego/v2/objects"
	"github.com/jackc/pgx/v5/pgtype"
	"server/bot/channels"
	"server/db"
	"strconv"
)

func SendMsgToChannel(bot *telego.Bot, update *objects.Update) {

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

	subscribedChannels, err := db.New(conn).ChannelsByAdmin(db.Ctx, pgtype.Int8{
		Int64: int64(update.Message.Chat.Id),
		Valid: true,
	})

	if len(subscribedChannels) <= 0 {
		bot.SendMessage(
			update.Message.Chat.Id,
			"You aren't admin of any channel",
			"",
			update.Message.MessageId,
			false,
			false)
		return
	}

	kb := bot.CreateInlineKeyboard()

	for i, subscriptionToGroupList := range subscribedChannels {
		grl, err := db.New(conn).ChannelById(db.Ctx, int64(subscriptionToGroupList.ChannelsID.Int32))

		if err != nil {
			bot.SendMessage(
				update.Message.Chat.Id,
				"Can't get channel from db",
				"",
				update.Message.MessageId,
				false,
				false)
			return
		}

		kb.AddCallbackButtonHandler(grl.Name, strconv.FormatInt(grl.ID, 10), i+1, func(u *objects.Update) {
			nextMsgChan, _ := bot.AdvancedMode().RegisterChannel(strconv.Itoa(u.CallbackQuery.Message.Chat.Id), "message")

			bot.SendMessage(
				update.Message.Chat.Id,
				"Enter your msg",
				"",
				update.Message.MessageId,
				false,
				false)

			nextUpdate := <-*nextMsgChan

			c, _ := db.NewConn()

			defer c.Close(db.Ctx)

			err := channels.SentMsg(bot, c, nextUpdate.Message.Text, grl)

			if err != nil {
				bot.SendMessage(
					nextUpdate.Message.Chat.Id,
					"Error with sending msg",
					"",
					nextUpdate.Message.MessageId,
					false,
					false)
				return
			}

			//groupListID, _ := strconv.ParseInt(u.CallbackQuery.Data, 10, 64)
			//chats, err := db.New(c).SubscriptionToChannelsByGroupListId(db.Ctx, pgtype.Int4{
			//	Valid: true,
			//	Int32: int32(groupListID),
			//})
			//
			//if err != nil {
			//	bot.SendMessage(
			//		nextUpdate.Message.Chat.Id,
			//		"Can't get chats from db",
			//		"",
			//		nextUpdate.Message.MessageId,
			//		false,
			//		false)
			//	return
			//}
			//
			//for _, chat := range chats {
			//	_, err = bot.SendMessage(
			//		int(chat.ChatID.Int64),
			//		fmt.Sprintf("%s: %s", grl.Name, nextUpdate.Message.Text),
			//		"",
			//		0,
			//		false,
			//		false)
			//
			//	if err != nil {
			//		log.Println(err.Error(), int(chat.ChatID.Int64))
			//	}
			//}

			bot.SendMessage(
				nextUpdate.Message.Chat.Id,
				"Sent",
				"",
				nextUpdate.Message.MessageId,
				false,
				false)
		})
	}

	_, err = bot.AdvancedMode().ASendMessage(update.Message.Chat.Id,
		"Select group which you want to send",
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
