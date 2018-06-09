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
	color := ""
	if member.Executive {
		memberPrivStr = "取締役メンバー"
		color = "#ED1A3D"
	} else {
		memberPrivStr = "一般メンバー"
		color = "#008000"
	}
	msg := "*" + member.Code + "*\n*" + member.Name + "*\n" + memberPrivStr + "\n" + member.Role + "担当\n\n" + member.Description

	if member.Secession {
		memberPrivStr += "(脱退済み)"
		color = "#000000"
	}

	attachment := slack.Attachment{
		Color: color,
		Text:  msg,
	}
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	b.client.PostMessage(ev.Channel, "", params)
	return nil
}
