package extapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/TinyKitten/sugoibot/constant"
)

// TokyoMetroODResponse 東京メトロオープンデータのレスポンス構造体
type TokyoMetroODResponse struct {
	Context              string    `json:"@context"`
	ID                   string    `json:"@id"`
	Date                 time.Time `json:"dc:date"`
	Valid                time.Time `json:"dct:valid"`
	Operator             string    `json:"odpt:operator"`
	Railway              string    `json:"odpt:railway"`
	TimeOfOrigin         time.Time `json:"odpt:timeOfOrigin"`
	TrainInformationText string    `json:"odpt:trainInformationText"`
	Type                 string    `json:"@type"`
}

// GetTrainDelayArray 東京メトロの遅延情報をすべて取得
func GetLineInformationArray() ([]TokyoMetroODResponse, error) {
	tokyoMetroToken := os.Getenv("TOKYOMETRO_OD_TOKEN")
	if tokyoMetroToken == "" {
		return []TokyoMetroODResponse{}, errors.New(constant.ERR_TOKYOMETRO_OD_TOKEN_NOT_DEFINED)
	}
	resp, err := http.Get("https://api.tokyometroapp.jp/api/v2/datapoints?rdf:type=odpt:TrainInformation&acl:consumerKey=" + tokyoMetroToken)
	if err != nil {
		return []TokyoMetroODResponse{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var data []TokyoMetroODResponse
	err = decoder.Decode(&data)
	if err != nil {
		return []TokyoMetroODResponse{}, err
	}

	return data, nil
}

// GetLineInformationByODPTRailway odpt:railwayで路線名を練り込む
func GetLineInformationByODPTRailway(odptRailway string) (TokyoMetroODResponse, error) {
	arr, err := GetLineInformationArray()
	if err != nil {
		return TokyoMetroODResponse{}, err
	}
	matchedLine := new(TokyoMetroODResponse)
	for _, line := range arr {
		if line.Railway == odptRailway {
			matchedLine = &line
		}
	}
	if matchedLine == nil {
		return TokyoMetroODResponse{}, errors.New(constant.ERR_TOKYOMETRO_LINE_NOT_FOUND)
	}
	return *matchedLine, nil
}

// ConvertODPTRailwayToJP odpt:railwayを日本語に変換する
func ConvertODPTRailwayToJP(odptRailway string) string {
	switch odptRailway {
	case "odpt.Railway:TokyoMetro.Hibiya":
		return "日比谷線"
	case "odpt.Railway:TokyoMetro.Chiyoda":
		return "千代田線"
	case "odpt.Railway:TokyoMetro.Tozai":
		return "東西線"
	case "odpt.Railway:TokyoMetro.Fukutoshin":
		return "副都心線"
	case "odpt.Railway:TokyoMetro.Ginza":
		return "銀座線"
	case "odpt.Railway:TokyoMetro.Yurakucho":
		return "有楽町線"
	case "odpt.Railway:TokyoMetro.Hanzomon":
		return "半蔵門線"
	case "odpt.Railway:TokyoMetro.Namboku":
		return "南北線"
	case "odpt.Railway:TokyoMetro.Marunouchi":
		return "丸ノ内線"
	default:
		return "不明な路線"
	}
}
