package bot

import (
	"net"

	"github.com/TinyKitten/sugoibot/env"
	"github.com/nlopes/slack"
)

func (b *Bot) handleDebugMyHost(ev *slack.MessageEvent) error {
	if !env.Debug() {
		failmsg := "や、デバッグモード有効になってNASA🚀"
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage(failmsg, ev.Channel))
		return nil
	}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		b.handleError(err, ev)
	}
	var ip string
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
			}
		}
	}
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(ip, ev.Channel))
	return nil
}
