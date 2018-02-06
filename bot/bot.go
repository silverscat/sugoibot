package bot

import (
	"bytes"
	"log"
	"strconv"
	"strings"

	"github.com/TinyKitten/sugoibot/env"
	"github.com/TinyKitten/sugoibot/extapi"
	"github.com/nlopes/slack"
)

// Bot Bot構造体
type Bot struct {
	client *slack.Client
	rtm    *slack.RTM
}

// NewBot Botの構造体を生成
func NewBot(token string) *Bot {
	client := slack.New(token)
	debug := env.Debug()
	client.SetDebug(debug)
	rtm := client.NewRTM()
	return &Bot{
		client: client,
		rtm:    rtm,
	}
}

// StartListenMessage メッセージの処理を開始する
func (b *Bot) StartListenMessage() {
	go b.rtm.ManageConnection()
	for msg := range b.rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := b.handleMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}

// handleMessageEvent メッセージを処理する
func (b *Bot) handleMessageEvent(ev *slack.MessageEvent) error {
	if ev.User == ev.BotID {
		return nil
	}
	cmdPrefixIndex := strings.Index(ev.Text, "./")
	if cmdPrefixIndex == -1 || cmdPrefixIndex != 0 {
		return nil
	}
	switch ev.Text {
	case "./dappun":
		b.handleDappun(ev)
		return nil
	case "./matsuya":
		b.handleMatsuya(ev)
		return nil
	case "./tokyometro_delay":
		b.handleTokyoMetroDelay(ev)
		return nil
	default:
		b.handleDefault(ev)
		return nil
	}
}

func (b *Bot) handleDappun(ev *slack.MessageEvent) error {
	butimili := "うおおおおおおおおおおおおあああああああああああああああああああああああああああああああ！！！！！！！！！！！ (ﾌﾞﾘﾌﾞﾘﾌﾞﾘﾌﾞﾘｭﾘｭﾘｭﾘｭﾘｭﾘｭ！！！！！！ﾌﾞﾂﾁﾁﾌﾞﾌﾞﾌﾞﾁﾁﾁﾁﾌﾞﾘﾘｲﾘﾌﾞﾌﾞﾌﾞﾌﾞｩｩｩｩｯｯｯ！！！！！！！)"
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(butimili, ev.Channel))
	return nil
}

func (b *Bot) handleDefault(ev *slack.MessageEvent) error {
	nasa := "や、そんなコマンドのNASA✋"
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(nasa, ev.Channel))
	return nil
}

func (b *Bot) handleMatsuya(ev *slack.MessageEvent) error {
	menu, err := extapi.GetRandom()
	if err != nil {
		return b.handleError(err, ev)
	}
	price := strconv.Itoa(menu.Price)
	resp := "*" + menu.Name + "*\n*" + price + "円*\n" + menu.Description + "\n" + menu.ImageURL
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(resp, ev.Channel))
	return nil
}

func (b *Bot) handleTokyoMetroDelay(ev *slack.MessageEvent) error {
	lines, err := extapi.GetLineInformationArray()
	if err != nil {
		return b.handleError(err, ev)
	}

	var buffer bytes.Buffer
	for i, line := range lines {
		lineName := extapi.ConvertODPTRailwayToJP(line.Railway)
		var innerBuffer bytes.Buffer
		innerBuffer.WriteString("*")
		innerBuffer.WriteString(lineName)
		innerBuffer.WriteString("*")
		innerBuffer.WriteString(" ")
		innerBuffer.WriteString(line.TrainInformationText)
		if i != len(lines) {
			innerBuffer.WriteString("\n")
		}
		buffer.WriteString(innerBuffer.String())
	}
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(buffer.String(), ev.Channel))
	return nil
}

func (b *Bot) handleError(err error, ev *slack.MessageEvent) error {
	log.Println(err)
	butimili := "エラーあああああああああああああああああああああああああああああああ！！！！！！！！！！！ (ﾌﾞﾘﾌﾞﾘﾌﾞﾘﾌﾞﾘｭﾘｭﾘｭﾘｭﾘｭﾘｭ！！！！！！ﾌﾞﾂﾁﾁﾌﾞﾌﾞﾌﾞﾁﾁﾁﾁﾌﾞﾘﾘｲﾘﾌﾞﾌﾞﾌﾞﾌﾞｩｩｩｩｯｯｯ！！！！！！！)"
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(butimili, ev.Channel))
	return err
}
