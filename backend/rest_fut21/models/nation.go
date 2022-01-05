package models

import "fmt"

type Nation struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (nation Nation) String() string {
	return fmt.Sprintf("{Id: %d, Name: %s", nation.Id, nation.Name)
}
