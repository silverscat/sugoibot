package extapi

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/TinyKitten/sugoibot/models"
)

// GetMemberByCode メンバーコードでメンバーを取得
func GetMemberByCode(code string) (*models.TKAPIMemberModel, error) {
	query := `
		{
		  member(code:"` + code + `") {
				code
				description
				executive
				id
				name
				role
				secession
				since
				twitter
				until
		  }
	  }
	  `
	url := "https://teamkitten-193615.appspot.com/"
	body, err := makeGraphQLRequest(url, query)
	if err != nil {
		return nil, err
	}
	resp := models.GraphQLResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	log.Println(query)
	if resp.Errors != nil {
		return nil, errors.New(resp.Errors[0].Message)
	}
	return resp.Data.Member, nil
}
