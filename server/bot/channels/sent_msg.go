package channels

import (
	"fmt"
	"github.com/SakoDroid/telego/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"log"
	"server/db"
)

func SentMsg(bot *telego.Bot, c *pgx.Conn, msg string, grl db.Channel) error {
	chats, err := db.New(c).SubscriptionToChannelsByChannelId(db.Ctx, pgtype.Int4{
		Valid: true,
		Int32: int32(grl.ID),
	})

	if err != nil {
		return err
	}

	for _, chat := range chats {
		_, err = bot.SendMessage(
			int(chat.ChatID.Int64),
			fmt.Sprintf("%s: %s", grl.Name, msg),
			"",
			0,
			false,
			false)

		if err != nil {
			log.Println(err.Error(), int(chat.ChatID.Int64))
		}
	}

	return nil
}
