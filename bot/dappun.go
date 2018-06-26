package bot

import (
	"math/rand"
	"time"

	"github.com/nlopes/slack"
)

func (b *Bot) handleDappun(ev *slack.MessageEvent, args ...string) error {
	butimili := []string{
		"*あああああああああああああああああああああああああああああああ!!!!!!!!!!!(ﾌﾞﾘﾌﾞﾘﾌﾞﾘﾌﾞﾘｭﾘｭﾘｭﾘｭﾘｭﾘｭ!!!!!!ﾌﾞﾂﾁﾁﾌﾞﾌﾞﾌﾞﾁﾁﾁﾁﾌﾞﾘﾘｲﾘﾌﾞﾌﾞﾌﾞﾌﾞｩｩｩｩｯｯｯ!!!!!!!)*",
		"*(ﾌﾞﾘﾌﾞﾘﾌﾞﾘﾌﾞﾘｭﾘｭﾘｭﾘｭﾘｭﾘｭ!!!!!!ﾌﾞﾂﾁﾁﾌﾞﾌﾞﾌﾞﾁﾁﾁﾁﾌﾞﾘﾘｲﾘﾌﾞﾌﾞﾌﾞﾌﾞｩｩｩｩｯｯｯ!!!!!!!)あああああああああああああああああああああああああああああああ!!!!!!!!!!!*",
		"*are are are are are are are are are!!!!!!!!!!!!!(Boo lee Boo lee Boo lee Boo Liu Liu Liu Liu Liu Liu !!!!!!!!!!!! Boo tea tea Boo Boo Boo tea tea tea tea Boo lee lee Boo Boo Boo Boo woo woo woo uh)*",
		"*うおおおおおおおおおおおおああああああああああああああああああ!!!!!!!!!!!(ﾌﾞﾘﾌﾞﾘﾌﾞﾘﾌﾞﾘｭﾘｭﾘｭﾘｭﾘｭﾘｭ!!!!!!ﾌﾞﾂﾁﾁﾌﾞﾌﾞﾌﾞﾁﾁﾁﾁﾌﾞﾘﾘｲﾘﾌﾞﾌﾞﾌﾞﾌﾞｩｩｩｩｯｯｯ!!!!!!!)*",
	}

	rand.Seed(time.Now().UnixNano())

	attachment := slack.Attachment{
		Color: "#994C00",
		Text:  butimili[rand.Intn(len(butimili)-1)],
	}
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	b.client.PostMessage(ev.Channel, "", params)
	return nil
}
