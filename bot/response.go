package bot

import "github.com/nlopes/slack"

func (b *Bot) handleCmdError(ev *slack.MessageEvent, msg string) {
	attachment := slack.Attachment{
		Color: "#ED1A3D",
		Text:  msg,
	}
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	b.client.PostMessage(ev.Channel, "", params)
}

func (b *Bot) handleCmdErrorWithPretext(ev *slack.MessageEvent, msg, pre string) {
	attachment := slack.Attachment{
		Color:   "#ED1A3D",
		Text:    msg,
		Pretext: pre,
	}
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	b.client.PostMessage(ev.Channel, "", params)
}

func (b *Bot) handleCmdCompleted(ev *slack.MessageEvent, msg string) {
	attachment := slack.Attachment{
		Color: "#008000",
		Text:  msg,
	}
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	b.client.PostMessage(ev.Channel, "", params)
}

func (b *Bot) handleCmdCompletedWithPretext(ev *slack.MessageEvent, msg, pre string) {
	attachment := slack.Attachment{
		Color:   "#008000",
		Text:    msg,
		Pretext: pre,
	}
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	b.client.PostMessage(ev.Channel, "", params)
}
