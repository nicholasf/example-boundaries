package main

import(
	"github.com/nicholasf/example/distributed/boundaries"
	"github.com/nicholasf/example/distributed/boundaries/factories"
	"github.com/gin-gonic/gin"
)

type HousesApp interface {
	Houses() boundaries.Houses
}

type housesApp struct {
	houses boundaries.Houses
}

func NewHousesApp() HousesApp {
//	users := factories.NewUsers()
	houses := factories.NewHouses(users)

	return &housesApp{
		houses: houses,
	}
}

func runHouses(port string) {

}