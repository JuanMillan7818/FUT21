package models

import "fmt"

type Club struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (club Club) String() string {
	return fmt.Sprintf("{Id: %d, Name: %s", club.Id, club.Name)
}
