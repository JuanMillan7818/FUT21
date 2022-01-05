package models

import "fmt"

// Struct
type Player struct {
	Id           string
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	PositionFull string `json:"positionFull"`
	Natal        Nation `json:"nation"`
	Club         Club   `json:"club"`
}

// Construct
func NewPlayer(id, firstName, lastName, positionFull string, nation Nation, club Club) *Player {
	return &Player{
		Id:           id,
		FirstName:    firstName,
		LastName:     lastName,
		PositionFull: positionFull,
		Natal:        nation,
		Club:         club,
	}
}

// To String
func (playerTemp *Player) String() string {
	return fmt.Sprintf(
		"{id: %s, firstName: %s, lastName: %s, positionFull: %s, nation: %s, club: %s",
		playerTemp.Id,
		playerTemp.FirstName,
		playerTemp.LastName,
		playerTemp.PositionFull,
		playerTemp.Natal.String(),
		playerTemp.Club)
}

// GETTERS
func (playerTemp *Player) GetId() string {
	return playerTemp.Id
}
func (playerTemp *Player) GetFirstName() string {
	return playerTemp.FirstName
}
func (playerTemp *Player) GetLastName() string {
	return playerTemp.LastName
}
func (playerTemp *Player) GetPositionFull() string {
	return playerTemp.PositionFull
}
func (playerTemp *Player) GetNation() Nation {
	return playerTemp.Natal
}
func (playerTemp *Player) GetClub() Club {
	return playerTemp.Club
}

// SETTERS
func (playerTemp *Player) SetId(idNew string) {
	playerTemp.Id = idNew
}
func (playerTemp *Player) SetFirstName(firstNew string) {
	playerTemp.FirstName = firstNew
}
func (playerTemp *Player) SetLastName(lastNew string) {
	playerTemp.LastName = lastNew
}
func (playerTemp *Player) SetPosition(positionNew string) {
	playerTemp.PositionFull = positionNew
}
func (playerTemp *Player) SetNation(nationNew Nation) {
	playerTemp.Natal = nationNew
}
func (playerTemp *Player) SetClub(clubNew Club) {
	playerTemp.Club = clubNew
}
