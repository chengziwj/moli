package misc

import (
	"regexp"
)

const defaultLimit = 10

var re, _ = regexp.Compile(`(?i:^select (.+?) from) *`)
var countStatement = []byte("select count(1) from ")

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Total  int `json:"total"`
	Sql    string
}

func (p Pagination) CountSql(sql string) string {
	rs := re.ReplaceAll([]byte(sql), countStatement)
	return string(rs)
}


func NewPagination(page, limit int) Pagination {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = defaultLimit
	}
	offset := (page - 1) * limit
	return Pagination{Page: page, Limit: limit, Offset: offset}
}
