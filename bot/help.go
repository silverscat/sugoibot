package bot

import "github.com/nlopes/slack"

func (b *Bot) handleHelp(ev *slack.MessageEvent) error {
	butimili := `./getMenberByCode [メンバーコード]
	メンバーコードからプロフィールを取得します。

	./getMemberByReply [対象]
	リプライからプロフィールを取得します。

	./delayline [路線名]
	路線の遅延を確認します。対応している路線は./delayline listで確認できます。

	./dappun
	ど、どうしたんだいきなり大声出して。

	./matsuya
	ランダムで松屋のメニューを表示します。

	./tokyometro_delay
	東京メトロの遅延を確認します。
	`
	b.rtm.SendMessage(b.rtm.NewOutgoingMessage(butimili, ev.Channel))
	return nil
}
