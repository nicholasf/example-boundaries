package main


import (
	"github.com/gin-gonic/gin"
	usersTransport "github.com/nicholasf/example/distributed/users/transport"
)

func MapRoutes(rg *gin.RouterGroup, users UsersApp) *gin.RouterGroup {
	rg = usersTransport.MapRoutes(rg, users)
	return rg
}
