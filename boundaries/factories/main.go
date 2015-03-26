package factories

import (
	"github.com/nicholasf/example/boundaries"
	"github.com/nicholasf/example/houses"
	"github.com/nicholasf/example/users"
)

func NewHouses() boundaries.Houses {
	return houses.NewHouses()
}

func NewUsers() boundaries.Users {
	return users.NewUsers()
}