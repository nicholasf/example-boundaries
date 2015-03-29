package main

import(
	"github.com/nicholasf/example/monolith/boundaries"
	"github.com/nicholasf/example/monolith/boundaries/factories"
	"github.com/gin-gonic/gin"
)

type ExampleApp interface {
	Houses() boundaries.Houses
	Users() boundaries.Users
}

type exampleApp struct {
	houses boundaries.Houses
	users boundaries.Users
}

func NewExampleApp() ExampleApp {
	users := factories.NewUsers()
	houses := factories.NewHouses(users)

	return &exampleApp{
		houses: houses,
		users: users,
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
