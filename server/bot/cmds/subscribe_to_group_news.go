package cmds

import (
	"fmt"
	"github.com/SakoDroid/telego/v2"
	"github.com/SakoDroid/telego/v2/objects"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"server/db"
	"server/helper"
	"slices"
	"strconv"
)

func SubscribeToGroupNews(bot *telego.Bot, update *objects.Update) {

	log.Println(update.Message.Chat.Id)

	chatManager := bot.GetChatManagerById(update.Message.Chat.Id)
	if update.Message.Chat.Type != "private" {
		admins, err := chatManager.GetAdmins()

		if err != nil || !admins.Ok {
			bot.SendMessage(
				update.Message.Chat.Id,
				"Can't get chat admins",
				"",
				update.Message.MessageId,
				false,
				false)
			return
		}

		index := slices.IndexFunc(admins.Result, func(owner objects.ChatMemberOwner) bool {
			return owner.User.Id == update.Message.From.Id
		})

		if index == -1 {
			bot.SendMessage(
				update.Message.Chat.Id,
				"It can do only admin",
				"",
				update.Message.MessageId,
				false,
				false)
			return
		}
	}

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

	channels, err := db.New(conn).ListChannels(db.Ctx)

	subscribedChannels, err := db.New(conn).SubscriptionToChannelsByChatId(db.Ctx, pgtype.Int8{
		Int64: int64(update.Message.Chat.Id),
		Valid: true,
	})

	kb := bot.CreateInlineKeyboard()

	for i, channel := range helper.FilterSlices(channels, func(glist db.Channel) bool {
		return slices.IndexFunc(subscribedChannels, func(stlist db.SubscriptionToChannel) bool {
			return stlist.ChannelsID.Int32 == int32(glist.ID)
		}) == -1
	}) {
		kb.AddCallbackButtonHandler(channel.Name, strconv.FormatInt(channel.ID, 10), i+1, func(u *objects.Update) {
			c, _ := db.NewConn()

			defer c.Close(db.Ctx)

			channelID, _ := strconv.ParseInt(u.CallbackQuery.Data, 10, 64)
			db.New(c).InsertSubscriptionToChannel(db.Ctx, db.InsertSubscriptionToChannelParams{
				ChatID: pgtype.Int8{
					Int64: int64(u.CallbackQuery.Message.Chat.Id),
					Valid: true,
				},
				ChannelsID: pgtype.Int4{
					Int32: int32(channelID),
					Valid: true,
				},
				Username: pgtype.Text{
					String: u.CallbackQuery.Message.Chat.Username,
					Valid:  true,
				},
				Title: pgtype.Text{
					String: u.CallbackQuery.Message.Chat.Title,
					Valid:  true,
				},
				ChatType: pgtype.Text{
					String: u.CallbackQuery.Message.Chat.Type,
					Valid:  true,
				},
			})

			bot.SendMessage(
				u.CallbackQuery.Message.Chat.Id,
				"Complete! You'll receive msgs for this", "",
				u.CallbackQuery.Message.MessageId,
				false, false)
		})
	}

	_, err = bot.AdvancedMode().ASendMessage(update.Message.Chat.Id,
		"Select group which you want to subscribe",
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
