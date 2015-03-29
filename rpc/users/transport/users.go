package transport

import (
	"github.com/gin-gonic/gin"
	"github.com/nicholasf/example/rpc/boundaries"
	"log"
)

func Resource(boundary boundaries.Users) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Params.ByName("name")
		resource, err := boundary.ByName(name)

		if err != nil {
			ctx.JSON(400, err)
			log.Print("Could not locate user! ", name,  err)
			return
		}

		log.Print("Found user! ", resource)
		ctx.JSON(200, resource)
	}
}
