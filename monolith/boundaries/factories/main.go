package factories

import (
	"github.com/nicholasf/example/monolith/boundaries"
	"github.com/nicholasf/example/monolith/houses"
	"github.com/nicholasf/example/monolith/users"
)

func NewHouses(users boundaries.Users) boundaries.Houses {
	return houses.NewHouses(users)
}

func NewUsers() boundaries.Users {
	return users.NewUsers()
}