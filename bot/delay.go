package bot

import (
	"bytes"
	"strings"

	"github.com/TinyKitten/sugoibot/extapi"
	"github.com/TinyKitten/sugoibot/utils"

	"github.com/nlopes/slack"
)

func (b *Bot) handleDelay(ev *slack.MessageEvent) error {
	if ev.Text == "./delayline list" {
		var buffer bytes.Buffer
		compatibleLines := utils.LoadCSV("./assets/compatibleLine.csv")
		for _, line := range compatibleLines {
			buffer.WriteString(line[0])
			buffer.WriteString("\n")
		}
		attachment := slack.Attachment{
			Text:    buffer.String(),
			Pretext: "こちらです:",
		}
		params := slack.PostMessageParameters{}
		params.Attachments = []slack.Attachment{attachment}
		b.client.PostMessage(ev.Channel, "", params)
		return nil
	}

	lineName := strings.Trim(ev.Text, "./delayline ")
	lineNameSpaces := strings.Fields(lineName)
	if len(lineNameSpaces) != 1 {
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage("路線を指定してください。", ev.Channel))
		return nil
	}

	compatible := false
	compatibleLines := utils.LoadCSV("./assets/compatibleLine.csv")
	for _, line := range compatibleLines {
		if lineName == line[0] {
			compatible = true
		}
	}

	if !compatible {
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage(lineName+"は対応していない路線です。", ev.Channel))
		return nil
	}

	delayed, err := extapi.GetDelayedLineByName(lineName)
	if err != nil {
		b.handleError(err, ev)
	}
	attachment := slack.Attachment{
		Color:   "#4CAF50",
		Text:    lineName,
		Pretext: "遅延していません。",
	}
	if delayed {
		attachment = slack.Attachment{
			Color:   "#F44336",
			Text:    lineName,
			Pretext: "遅延しています！",
		}
	}
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	b.client.PostMessage(ev.Channel, "", params)
	return nil
}
