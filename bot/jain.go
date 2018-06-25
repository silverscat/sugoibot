package bot

import (
	"math/rand"
	"strings"

	"github.com/nlopes/slack"
)

func (b *Bot) handleJain(ev *slack.MessageEvent, args ...string) error {
	jain := `「邪淫」について‥‥
	お話します‥‥
	みんな‥‥
	「邪淫」って‥‥
	知ってるかな？
	「邪淫」というのはね
	たとえば
	おしっこするところを‥‥
	さわると
	「気持ちがいい」　とか
	あるいは‥‥
	おしっこするところを‥‥
	こすりつけると
	「気持ちがいい」
	といったことを
	「邪淫」というんだ。
	よい子のみんなの‥‥
	こころは‥‥
	心臓に
	あるんだよ。
	そして、心臓と
	おしっこするところ‥‥
	どちらが‥‥
	上かな？‥‥
	もちろん、‥‥
	心臓のほうが‥‥
	頭に近いから‥‥
	上だよね？‥‥
	からだの‥‥
	下の部分に	
	心が
	集中するとね‥‥
	そのぉ子は
	下の世界に‥‥
	生まれ変わるんだって。
	イヤだねぇ。
	今
	「邪淫」を‥‥
	行っていない子は‥‥
	これから先‥‥
	「邪淫」を
	しないようにしようね。
	今
	「邪淫」を行っている
	良い子は、
	やめようね！
	そして、‥‥
	お父さん
	お母さんを含めた
	みんなを
	大好きになって、
	みんなのために
	生きようね！`

	jainSlice := strings.Split(jain, "\n")
	n := rand.Intn(len(jainSlice))

	attachment := slack.Attachment{
		Color: "#994C00",
		Text:  jainSlice[n],
	}
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	b.client.PostMessage(ev.Channel, "", params)
	return nil
}
