package bot

import (
	"strconv"

	"github.com/TinyKitten/sugoibot/extapi"
	"github.com/nlopes/slack"
)

func (b *Bot) handleMatsuya(ev *slack.MessageEvent, args ...string) error {
	menu, err := extapi.GetRandom()
	if err != nil {
		return b.handleError(err, ev)
	}
	price := strconv.Itoa(menu.Price)
	resp := "*" + menu.Name + "*\n*" + price + "円*\n" + menu.Description + "\n" + menu.ImageURL
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(resp, ev.Channel))
	return nil
}
