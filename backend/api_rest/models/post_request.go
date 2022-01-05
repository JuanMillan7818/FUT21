package models

type TeamRequest struct {
	Name string `json:"Name"`
	Page int    `json:"Page"`
}

type TeamResponse struct {
	Page       int `json:"Page"`
	TotalPages int `json:"TotalPages"`
	Items      int `json:"Items"`
	TotalItems int `json:"TotalItems"`
	Players    []Player
}

type Player struct {
	Id           string `json:"player_id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PositionFull string `json:"positionFull"`
	Natal        string `json:"nation"`
	Club         string `json:"club"`
}
