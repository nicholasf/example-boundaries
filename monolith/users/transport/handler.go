package transport

import (
	"github.com/gin-gonic/gin"

	"github.com/nicholasf/example/monolith/boundaries"
)

func MapRoutes(r *gin.RouterGroup, u boundaries.Users) *gin.RouterGroup {
	r.GET("/user/:name", Resource(u))

	return r
}


