package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/nicholasf/example/boundaries"
)

func Resource(boundary boundaries.Users) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Params.ByName("name")

		user, _ := boundary.ByName(name)
		ctx.JSON(200, user)
	}
}
