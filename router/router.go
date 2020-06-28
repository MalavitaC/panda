package router

import (
	"panda/controller/integral"
	"panda/controller/user"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", user.Query)
		api.GET("/integrals", integral.Query)
	}
}
