package bot

import (
	"math/rand"
	"strings"

	"github.com/nlopes/slack"
)

func (b *Bot) handleOkayama(ev *slack.MessageEvent, args ...string) error {
	dovaahland := `昨日の8月15日にいつもの浮浪者のおっさん（60歳）と先日メールくれた汚れ好きの土方のにいちゃん
	（45歳）とわし（53歳）の3人で県北にある川の土手の下で盛りあったぜ。
	今日は明日が休みなんでコンビニで酒とつまみを買ってから滅多に人が来ない所なんで、
	そこでしこたま酒を飲んでからやりはじめたんや。
	3人でちんぽ舐めあいながら地下足袋だけになり持って来たいちぢく浣腸を3本ずつ入れあった。
	しばらくしたら、けつの穴がひくひくして来るし、糞が出口を求めて腹の中でぐるぐるしている。
	浮浪者のおっさんにけつの穴をなめさせながら、兄ちゃんのけつの穴を舐めてたら、
	先に兄ちゃんがわしの口に糞をドバーっと出して来た。
	それと同時におっさんもわしも糞を出したんや。もう顔中、糞まみれや、
	3人で出した糞を手で掬いながらお互いの体にぬりあったり、
	糞まみれのちんぽを舐めあって小便で浣腸したりした。ああ～～たまらねえぜ。
	しばらくやりまくってから又浣腸をしあうともう気が狂う程気持ちええんじゃ。
	浮浪者のおっさんのけつの穴にわしのちんぽを突うずるっ込んでやると
	けつの穴が糞と小便でずるずるして気持ちが良い。
	にいちゃんもおっさんの口にちんぽ突っ込んで腰をつかって居る。
	糞まみれのおっさんのちんぽを掻きながら、思い切り射精したんや。
	それからは、もうめちゃくちゃにおっさんと兄ちゃんの糞ちんぽを舐めあい、
	糞を塗りあい、二回も男汁を出した。もう一度やりたいぜ。
	やはり大勢で糞まみれになると最高やで。こんな、変態親父と糞あそびしないか。
	ああ～～早く糞まみれになろうぜ。
	岡山の県北であえる奴なら最高や。わしは163*90*53,おっさんは165*75*60、や
	糞まみれでやりたいやつ、至急、メールくれや。
	土方姿のまま浣腸して、糞だらけでやろうや。`

	dovaahlandSlice := strings.Split(dovaahland, "\n")
	n := rand.Intn(len(dovaahlandSlice))

	attachment := slack.Attachment{
		Color: "#994C00",
		Text:  dovaahlandSlice[n],
	}
	params := slack.PostMessageParameters{}
	params.Attachments = []slack.Attachment{attachment}
	b.client.PostMessage(ev.Channel, "", params)
	return nil
}
