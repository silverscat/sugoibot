package bot

import "github.com/nlopes/slack"

func (b *Bot) handleHelp(ev *slack.MessageEvent) error {
	msg := `
	*./getMenberByCode* _[メンバーコード]_
	メンバーコードからプロフィールを取得します。

	*./getMemberByReply* _[対象]_
	リプライからプロフィールを取得します。

	*./delayline* _[路線名]_
	路線の遅延を確認します。対応している路線は *./delayline list* で確認できます。

	*./dappun*
	ど、どうしたんだいきなり大声出して。

	*./matsuya*
	ランダムで松屋のメニューを表示します。

	*./tokyometro_delay*	
	東京メトロの遅延を確認します。
	`
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(msg, ev.Channel))
	return nil
}
