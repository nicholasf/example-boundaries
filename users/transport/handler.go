package transport

import (
	"github.com/gin-gonic/gin"

	"github.com/nicholasf/example/boundaries"
)

func MapRoutes(r *gin.RouterGroup, u boundaries.Users) *gin.RouterGroup {
	r.GET("/house", Resource(u))

	return r
}


