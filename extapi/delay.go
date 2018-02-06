package extapi

import (
	"encoding/json"
	"net/http"
)

// DelayResponse 遅延APIの構造体
type DelayResponse struct {
	Name          string `json:"name"`
	Company       string `json:"company"`
	LastUpdateGMT int    `json:"lastupdate_gmt"`
	Source        string `json:"source"`
}

// GetDelayedLineArray 遅延している路線の配列
func GetDelayedLineArray() ([]DelayResponse, error) {
	resp, err := http.Get("https://rti-giken.jp/fhc/api/train_tetsudo/delay.json")
	if err != nil {
		return []DelayResponse{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var data []DelayResponse
	err = decoder.Decode(&data)
	if err != nil {
		return []DelayResponse{}, err
	}

	return data, nil
}

// GetDelayedLineByName 遅延している路線を名前から
func GetDelayedLineByName(name string) (bool, error) {
	delayedLines, err := GetDelayedLineArray()
	if err != nil {
		return false, nil
	}
	delayed := false
	for _, line := range delayedLines {
		if line.Name == name {
			delayed = true
		}
	}
	return delayed, nil
}
