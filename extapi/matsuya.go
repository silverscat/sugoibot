package extapi

import (
	"encoding/json"
	"net/http"
)

// MatsuyaV4Response Matsuya-Web-APIのv4レスポンス
type MatsuyaV4Response struct {
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Price          int     `json:"price"`
	Calorie        int     `json:"calorie"`
	Protein        float32 `json:"protein"`
	Lipid          float32 `json:"lipid"`
	Carbohydrate   float32 `json:"carbohydrate"`
	Sodium         int     `json:"sodium"`
	SaltEquivalent float32 `json:"saltEquivalent"`
	Description    string  `json:"description"`
	ImageURL       string  `json:"imageURL"`
}

// GetRandom ランダムでメニューを取得
func GetRandom() (MatsuyaV4Response, error) {
	resp, err := http.Get("https://matsuya.makotia.me/v4/random")
	if err != nil {
		return MatsuyaV4Response{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var data MatsuyaV4Response
	err = decoder.Decode(&data)
	if err != nil {
		return MatsuyaV4Response{}, err
	}

	return data, nil
}
