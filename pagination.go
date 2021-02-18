package moli

const defaultLimit = 10

type Pagination struct {
	Limit  int
	Offset int
	Page   int
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
