package models

type Query struct {
	Page    string `json:"page"`
	PageNum string `json:"pageNum"`
	Order   string `json:"order"`
}
