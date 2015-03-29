package factories

import (
	"github.com/nicholasf/example/rpc/boundaries"
	"github.com/nicholasf/example/rpc/houses"
	"github.com/nicholasf/example/rpc/users/users"
)

func NewHouses(users boundaries.Users) boundaries.Houses {
	return houses.NewHouses(users)
}

func NewUsers(serviceAddress string) boundaries.Users {
	return users.NewUsers(serviceAddress)
}
