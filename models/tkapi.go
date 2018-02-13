package models

import "time"

// TKAPIMemberModel TeamKitten APIのメンバーモデル
type TKAPIMemberModel struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Executive   bool      `json:"executive"`
	Name        string    `json:"name"`
	Seccession  bool      `json:"seccession"`
	Since       time.Time `json:"since"`
	Twitter     string    `json:"twitter"`
	Role        string    `json:"role"`
	Until       time.Time `json:"until"`
}

// GraphQLResponse TeamKitten APIのメンバーレスポンス
type GraphQLResponse struct {
	Data   *GraphQLData    `json:"data,omitempty"`
	Errors []*GraphQLError `json:"errors,omitempty"`
}

type GraphQLData struct {
	Member *TKAPIMemberModel `json:"member,omitempty"`
}

// GraphQLError GraphQLエラー
type GraphQLError struct {
	Message   string                   `json:"message"`
	Locations []*GraphQLErrorLocations `json:"locations"`
}

// GraphQLErrorLocations GraphQLのエラー位置
type GraphQLErrorLocations struct {
	Line int `json:"line"`
	Row  int `json:"column"`
}
