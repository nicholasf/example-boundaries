package factories

import (
	"github.com/nicholasf/example/boundaries"
	"github.com/nicholasf/example/houses"
	"github.com/nicholasf/example/users"
)

func NewHouses(users boundaries.Users) boundaries.Houses {
	return houses.NewHouses(users)
}

func NewUsers() boundaries.Users {
	return users.NewUsers()
}