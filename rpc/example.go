package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nicholasf/example/rpc/boundaries"
	"github.com/nicholasf/example/rpc/boundaries/factories"
)

type ExampleApp interface {
	Houses() boundaries.Houses
	Users() boundaries.Users
}

type exampleApp struct {
	houses boundaries.Houses
	users  boundaries.Users
}

func NewExampleApp() ExampleApp {
	users := factories.NewUsers(":3000")
	houses := factories.NewHouses(users)

	return &exampleApp{
		houses: houses,
		users:  users,
	}
}

func (f *exampleApp) Houses() boundaries.Houses {
	return f.houses
}

func (f *exampleApp) Users() boundaries.Users {
	return f.users
}

func main() {
	example := NewExampleApp()

	err := example.Users().CreateUser("Joe")

	if err != nil {
		panic(err)
	}

	err = example.Houses().CreateHouse("21 Jump Street", "Joe")

	if err != nil {
		panic(err)
	}

	r := gin.New()
	rg := r.Group("/")
	rg = MapRoutes(rg, example)

	r.Run(":8000")
}
