package main


import (
	"github.com/gin-gonic/gin"
	usersTransport "github.com/nicholasf/example/users/transport"
	housesTransport "github.com/nicholasf/example/houses/transport"
)

func MapRoutes(rg *gin.RouterGroup, ex ExampleApp) *gin.RouterGroup {
	rg = usersTransport.MapRoutes(rg, ex.Users())
	rg = housesTransport.MapRoutes(rg, ex.Houses())
	return rg
}
