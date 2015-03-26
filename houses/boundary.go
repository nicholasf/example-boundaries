package houses

import (
	"github.com/nicholasf/example/houses/context/models"
)

type Houses struct {}

func (u *Houses) ByAddress(name string) (models.House, error) {
	return
}

func NewHouses() *Houses {
	return &Houses{}
}