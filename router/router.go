package router

import (
	"panda/controller/integral"
	"panda/controller/user"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	api := r.Group("/puzzle/api")
	{
		api.GET("/users", user.Query)
		api.POST("/users/sync", user.SyncUserInfo)
		api.POST("/users/login", user.Login)
		api.GET("/integrals", integral.Query)
	}
}
