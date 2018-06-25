package bot

import (
	"log"
	"strings"

	"github.com/TinyKitten/sugoibot/env"
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

	splitted := strings.Fields(ev.Text)

	switch splitted[0] {
	case "./getMemberByCode":
		err := b.handleGetMemberByCode(ev, splitted[1:]...)
		if err != nil {
			b.handleError(err, ev)
		}
		return err
	case "./getMemberByReply":
		err := b.handleGetMemberByReply(ev, splitted[1:]...)
		if err != nil {
			b.handleError(err, ev)
		}
		return err
	case "./todo":
		err := b.handleTodo(ev, splitted[1:]...)
		if err != nil {
			b.handleError(err, ev)
		}
		return err
	case "./dappun":
		return b.handleDappun(ev, splitted[1:]...)
	case "./matsuya":
		return b.handleMatsuya(ev, splitted[1:]...)
	case "./help":
		return b.handleHelp(ev, splitted[1:]...)
	case "./okayama":
		return b.handleOkayama(ev, splitted[1:]...)
	case "./jain":
		return b.handleJain(ev, splitted[1:]...)
	default:
		return b.handleDefault(ev)
	}
}

func (b *Bot) handleDefault(ev *slack.MessageEvent) error {
	nasa := "や、そんなコマンドのNASA✋"
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(nasa, ev.Channel))
	return nil
}

func (b *Bot) handleError(err error, ev *slack.MessageEvent) error {
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(err.Error(), ev.Channel))
	return err
}

func hasArgument(index int) bool {
	if index != -1 || index == 0 {
		return true
	}
	return false
}
