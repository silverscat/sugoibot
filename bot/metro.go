package bot

import (
	"bytes"

	"github.com/TinyKitten/sugoibot/extapi"
	"github.com/nlopes/slack"
)

func (b *Bot) handleTokyoMetroDelay(ev *slack.MessageEvent) error {
	lines, err := extapi.GetLineInformationArray()
	if err != nil {
		return b.handleError(err, ev)
	}

	var buffer bytes.Buffer
	delayed := false
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

		// 遅延しているか
		if line.TrainInformationText != "現在、平常どおり運転しています。" {
			delayed = true
		}
	}
	attachment := slack.Attachment{
		Color:   "#4CAF50",
		Text:    buffer.String(),
		Pretext: "遅延していません。",
	}
	if delayed {
		attachment = slack.Attachment{
			Color:   "#F44336",
			Text:    buffer.String(),
			Pretext: "一部遅延しています！",
		}
	}
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	b.client.PostMessage(ev.Channel, "", params)
	return nil
}
