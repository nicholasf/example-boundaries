package main

import (
	"github.com/gin-gonic/gin"
	housesTransport "github.com/nicholasf/example/rpc/houses/transport"
	usersTransport "github.com/nicholasf/example/rpc/users/transport"
)

func MapRoutes(rg *gin.RouterGroup, ex ExampleApp) *gin.RouterGroup {
	rg = usersTransport.MapRoutes(rg, ex.Users())
	rg = housesTransport.MapRoutes(rg, ex.Houses())
	return rg
}
