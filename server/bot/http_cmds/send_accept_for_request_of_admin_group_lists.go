package http_cmds

import (
	"fmt"
	"server/bot"
	"server/db"
)

func SendAcceptOfRequestOfAdminGroupLists(chatId int64, list db.GroupList) {
	bot.Bot.SendMessage(
		int(chatId),
		fmt.Sprintf("Your request for admins of %s was accept", list.Name),
		"",
		0,
		false,
		false)
}
