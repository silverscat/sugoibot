package extapi

import (
	"encoding/json"
	"errors"

	"github.com/TinyKitten/sugoibot/constant"
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
	url := "https://api.teamkitten.tk/"
	body, err := makeGraphQLRequest(url, query)
	if err != nil {
		return nil, err
	}
	resp := models.GraphQLMemberResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Errors != nil {
		return nil, errors.New(resp.Errors[0].Message)
	}
	return resp.Data.Member, nil
}

// GetMemberBySlackID SlackIDでメンバーを検索
func GetMemberBySlackID(id string) (*models.TKAPIMemberModel, error) {
	query := `
		{
		  members {
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
				slack
		  }
	  }
	  `
	url := "https://api.teamkitten.tk/"
	body, err := makeGraphQLRequest(url, query)
	if err != nil {
		return nil, err
	}
	resp := models.GraphQLMembersResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Errors != nil {
		return nil, errors.New(resp.Errors[0].Message)
	}

	for _, m := range resp.Data.Members {
		if m.Slack == id {
			return m, nil
		}
	}

	return nil, errors.New(constant.ERR_MEMBER_NOT_FOUND)
}
