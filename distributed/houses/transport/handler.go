package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/nicholasf/example/distributed/boundaries"
)

func MapRoutes(r *gin.RouterGroup, u boundaries.Houses) *gin.RouterGroup {
	r.GET("/house/:address", Resource(u))

	return r
}
