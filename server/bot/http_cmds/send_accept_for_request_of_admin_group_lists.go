package http_cmds

import (
	"fmt"
	"github.com/jackc/pgx/v5"
	"server/bot"
	"server/bot/channels"
	"server/db"
)

func SendAcceptOfRequestOfAdminChannels(chatId int64, list db.Channel) {
	bot.Bot.SendMessage(
		int(chatId),
		fmt.Sprintf("Your request for admins of %s was accept", list.Name),
		"",
		0,
		false,
		false)
}

func SendMsgToChannel(c *pgx.Conn, msg string, grl db.Channel) error {
	return channels.SentMsg(bot.Bot, c, msg, grl)
}
