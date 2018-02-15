package bot

import (
	"strings"

	"github.com/TinyKitten/sugoibot/extapi"

	"github.com/nlopes/slack"
)

func (b *Bot) handleGetMemberByCode(ev *slack.MessageEvent) error {
	code := strings.Replace(ev.Text, "./getMemberByCode ", "", 1)
	codeSpaces := strings.Fields(code)
	if len(codeSpaces) != 1 {
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage("メンバーコードを指定してください。", ev.Channel))
		return nil
	}

	member, err := extapi.GetMemberByCode(code)
	if err != nil {
		return err
	}
	memberPrivStr := ""
	if member.Executive {
		memberPrivStr = "取締役メンバー"
	} else {
		memberPrivStr = "一般メンバー"
	}
	msg := code + ":\n" + member.Role + "担当\n" + member.Name + "\n" + memberPrivStr

	if member.Seccession {
		msg += "\n脱退済みメンバー"
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage(msg, ev.Channel))
		return nil
	}
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(msg, ev.Channel))
	return nil
}
