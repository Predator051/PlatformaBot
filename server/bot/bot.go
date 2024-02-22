package bot

import (
	"fmt"
	"github.com/SakoDroid/telego/v2"
	cfg "github.com/SakoDroid/telego/v2/configs"
	"github.com/SakoDroid/telego/v2/objects"
	"os"
	"server/bot/cmds"
)

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

	bot.AddHandler(string(cmds.RequestForAdminsPrivilegesName), func(update *objects.Update) {
		cmds.RequestForAdminsPrivileges(bot, update)
	}, "private")

	bot.Run(true)
	//
	//updates, _ := bot.UpdatesViaLongPolling(nil)
	//bh, err := th.NewBotHandler(bot, updates)
	//
	//if err != nil {
	//	println(err)
	//}
	//
	//defer bh.Stop()
	//defer bot.StopLongPolling()
	//
	//bh.Handle(func(bot *telego.Bot, update telego.Update) {
	//	println("Handle " + cmds.RequestForAdminsPrivilegesName)
	//	cmds.RequestForAdminsPrivileges(bot, update, bh)
	//}, th.CommandEqual(string(cmds.RequestForAdminsPrivilegesName)))
	//
	//bh.Start()
	//println("Stop bot!")
}
