package moli

const defaultLimit = 10

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Page   int `json:"page"`
	Total  int `json:"total"`
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
