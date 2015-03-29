package houses

import (
	"github.com/nicholasf/example/distributed/houses/context/models"
	"github.com/nicholasf/example/distributed/boundaries"
	"errors"
	"fmt"
	"log"
)

type Houses struct {
	Users boundaries.Users
	houses map[string]*models.House
}

func (h *Houses) ByAddress(address string) (house *models.House, err error) {
	log.Printf(fmt.Sprintf("Locating house by address [%s]", address))
	house, ok := h.houses[address]

	if !ok {
		err = errors.New("No such house")
	}

	return
}

func (h *Houses) CreateHouse(address, ownerName string) (err error) {
	_, ok := h.houses[address]

	if ok {
		err = errors.New("House already exists with address [" + address + "].")
		return
	}

	//dependency usage
	owner, err := h.Users.ByName(ownerName)

	if err != nil {
		err = errors.New("Unknown owner [" + ownerName + "].")
		return
	}

	h.houses[address] = &models.House{ Address: address, Owner: owner }
	return

}

func NewHouses(users boundaries.Users) *Houses {
	return &Houses{
		Users: users,
		houses: make(map[string]*models.House),
	}
}