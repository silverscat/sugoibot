package models

import "time"

// TKAPIMemberModel TeamKitten APIのメンバーモデル
type TKAPIMemberModel struct {
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Executive   bool      `json:"executive"`
	Name        string    `json:"name"`
	Secession   bool      `json:"secession"`
	Since       time.Time `json:"since"`
	Twitter     string    `json:"twitter"`
	Role        string    `json:"role"`
	Until       time.Time `json:"until"`
	Slack       string    `json:"slack"`
}
