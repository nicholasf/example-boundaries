package boundaries

import (
	userModels "github.com/nicholasf/example/users/context/models"
	houseModels "github.com/nicholasf/example/houses/context/models"
)

type Houses interface {
	ByAddress(name string) (*houseModels.House, error)
	CreateHouse(address, ownerName string) error

}

type Users interface {
	ByName(name string) (*userModels.User, error)
	CreateUser(name string) error
}