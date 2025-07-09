package models

type Page struct {
	Content       []interface{} `json:"data"`
	Total      int64         `json:"total"`
	Page       int64         `json:"page"`
	Limit      int64         `json:"limit"`
	SortBy     string        `json:"sort_by"`
	SortOrder  string        `json:"sort_order"`
	TotalPages int64         `json:"total_pages"`
}
