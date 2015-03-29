package factories

import (
	"github.com/nicholasf/example/distributed/boundaries"
	"github.com/nicholasf/example/distributed/houses"
	"github.com/nicholasf/example/distributed/users"
)

func NewHouses(users boundaries.Users) boundaries.Houses {
	return houses.NewHouses(users)
}

func NewUsers() boundaries.Users {
	return users.NewUsers()
}