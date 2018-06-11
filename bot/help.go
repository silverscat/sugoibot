package bot

import "github.com/nlopes/slack"

func (b *Bot) handleHelp(ev *slack.MessageEvent) error {
	msg := `
	*./getMenberByCode* _[メンバーコード]_
	メンバーコードからプロフィールを取得します。

	*./getMemberByReply* _[対象]_
	リプライからプロフィールを取得します。

	*./dappun*
	ど、どうしたんだいきなり大声出して。

	*./matsuya*
	ランダムで松屋のメニューを表示します。
	`
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(msg, ev.Channel))
	return nil
}
