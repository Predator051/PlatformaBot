package bot

import (
	"fmt"
	"github.com/SakoDroid/telego/v2"
	cfg "github.com/SakoDroid/telego/v2/configs"
	"github.com/SakoDroid/telego/v2/objects"
	"os"
	"server/bot/cmds"
	"slices"
	"strings"
)

var Bot *telego.Bot

func New(token string) {
	bot, err := telego.NewBot(&cfg.BotConfigs{
		BotAPI:         cfg.DefaultBotAPI,
		APIKey:         token,
		UpdateConfigs:  cfg.DefaultUpdateConfigs(),
		Webhook:        false,
		LogFileAddress: cfg.DefaultLogFile,
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Bot = bot

	bot.AddHandler(string(cmds.RequestForAdminsPrivilegesName), func(update *objects.Update) {
		cmds.RequestForAdminsPrivileges(bot, update)
	}, "private")

	updateChannel := *(bot.GetUpdateChannel())

	go func() {
		for {
			update := <-updateChannel

			if update.Message != nil {
				println(update.Message.Text)
			}

			if slices.Contains([]string{
				"private", "group",
			}, update.Message.Chat.Type) && strings.Contains(update.Message.Text, string(cmds.SubscribeToGroupNewsName)) {
				cmds.SubscribeToGroupNews(bot, update)
			}

			if strings.Contains(update.Message.Text, string(cmds.SendMsgToChannelName)) {
				cmds.SendMsgToChannel(bot, update)
			}
		}
	}()

	//bot.AddHandler(string(cmds.SubscribeToGroupNewsName), func(update *objects.Update) {
	//	cmds.SubscribeToGroupNews(bot, update)
	//}, "private", "group")

	bot.Run(true)
}
