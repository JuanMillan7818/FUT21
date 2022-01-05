package models

type Fut21 struct {
	List        []Player `json:"items"`
	Pages       int      `json:"totalPages"`
	TotalResult int      `json:"totalResults"`
	Items       int      `json:"count"`
}
