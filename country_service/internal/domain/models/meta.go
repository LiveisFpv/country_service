package models

type Filter struct {
	Field string
	Value string
}

type OrderBy struct {
	Field     string
	Direction string
}

type Pagination struct {
	Current int
	Total   int
	Limit   int
}
