package main

import(
	"github.com/nicholasf/example/distributed/boundaries"
	"github.com/nicholasf/example/distributed/boundaries/factories"
	"github.com/gin-gonic/gin"
)

type UsersApp interface {
	Users() boundaries.Houses
}

type usersApp struct {
	users boundaries.Users
}

func NewUsersApp() HousesApp {
	users := factories.NewUsers()
//	houses := factories.NewHouses(users)

	return &housesApp{
		users: users,
	}
}

//func main() {
//	example := NewHousesApp()
//
//	//	err := example.Users().CreateUser("Joe")
//
//	//	if err != nil {
//	//		panic(err)
//	//	}
//
//	err := example.Houses().CreateHouse("21 Jump Street", "Joe")
//
//	if err != nil {
//		panic(err)
//	}
//
//	r := gin.New()
//	rg := r.Group("/")
//	rg = MapRoutes(rg, example)
//
//	r.Run(":8000")
//}
