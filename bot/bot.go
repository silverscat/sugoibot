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
	// 引数が必要なコマンド
	// 管理者専用
	if ev.User == env.GetAdminID() {
		addUserIndex := strings.Index(ev.Text, "./adduser ")
		if hasArgument(addUserIndex) {
			err := b.handleAddUserAdmin(ev)
			if err != nil {
				b.handleError(err, ev)
			}
			return err
		}
	} else {
		// fallback
		addUserIndex := strings.Index(ev.Text, "./adduser ")
		if hasArgument(addUserIndex) {
			b.rtm.SendMessage(b.rtm.NewOutgoingMessage("実行権限がありません。", ev.Channel))
			return nil
		}
	}
	getMemberByCodeIndex := strings.Index(ev.Text, "./getMemberByCode ")
	if hasArgument(getMemberByCodeIndex) {
		err := b.handleGetMemberByCode(ev)
		if err != nil {
			b.handleError(err, ev)
		}
		return err
	}
	getMemberByReplyIndex := strings.Index(ev.Text, "./getMemberByReply ")
	if hasArgument(getMemberByReplyIndex) {
		err := b.handleGetMemberByReply(ev)
		if err != nil {
			b.handleError(err, ev)
		}
		return err
	}
	// 大抵の路線の遅延
	delayIndex := strings.Index(ev.Text, "./delayline ")
	if hasArgument(delayIndex) {
		return b.handleDelay(ev)
	}
	// 引数が必要ないコマンド
	// 脱糞
	if strings.HasPrefix(ev.Text, "./dappun") {
		return b.handleDappun(ev)
	}
	// 松屋
	if strings.HasPrefix(ev.Text, "./matsuya") {
		return b.handleMatsuya(ev)
	}
	// 東京メトロ遅延
	if strings.HasPrefix(ev.Text, "./tokyometro_delay") {
		return b.handleTokyoMetroDelay(ev)
	}
	// ヘルプ
	if strings.HasPrefix(ev.Text, "./help") {
		return b.handleHelp(ev)
	}
	// デバッグ用 ホスト名を吐く
	if strings.HasPrefix(ev.Text, "./debug_myip") {
		return b.handleDebugMyHost(ev)
	}
	return b.handleDefault(ev)
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
