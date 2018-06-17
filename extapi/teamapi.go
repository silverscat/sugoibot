package extapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/TinyKitten/sugoibot/constant"
	"github.com/TinyKitten/sugoibot/models"
)

var (
	// API すごいTeamKittenのすごいAPIのURL
	API = "https://api.teamkitten.tk/v1/members" // TODO: 戻す
)

// GetMemberByCode メンバーコードでメンバーを取得
func GetMemberByCode(code string) (*models.TKAPIMemberModel, error) {
	data, err := http.Get(API + "/" + code)
	if err != nil {
		return nil, err
	}
	defer data.Body.Close()
	if data.StatusCode == 404 {
		return nil, constant.ERR_MEMBER_NOT_FOUND
	}
	if data.StatusCode != 404 && data.StatusCode != 200 {
		return nil, constant.ERR_API_ERROR
	}
	member := &models.TKAPIMemberModel{}
	byteArray, _ := ioutil.ReadAll(data.Body)
	err = json.Unmarshal(byteArray, &member)
	if err != nil {
		return nil, err
	}
	return member, nil
}

// GetMemberBySlackID SlackIDでメンバーを検索
func GetMemberBySlackID(id string) (*models.TKAPIMemberModel, error) {
	data, err := http.Get(API)
	if err != nil {
		return nil, err
	}
	defer data.Body.Close()
	members := []*models.TKAPIMemberModel{}
	byteArray, _ := ioutil.ReadAll(data.Body)
	err = json.Unmarshal(byteArray, &members)
	if err != nil {
		return nil, err
	}

	for _, m := range members {
		if m.Slack == id {
			return m, nil
		}
	}

	return nil, constant.ERR_MEMBER_NOT_FOUND
}
