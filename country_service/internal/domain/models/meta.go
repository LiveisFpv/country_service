package models

type Pagination struct {
	Current int
	Total   int
	Limit   int
}

type Filter struct {
	Field string
	Value string
}

type Sort struct {
	Direction string
	By        string
}
