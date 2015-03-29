package main

import (
	"github.com/gin-gonic/gin"
	housesTransport "github.com/nicholasf/example/distributed/houses/transport"
)

func MapHouseRoutes(rg *gin.RouterGroup, houses HousesApp) *gin.RouterGroup {
	rg = housesTransport.MapRoutes(rg, ex.Houses())
	return rg
}
