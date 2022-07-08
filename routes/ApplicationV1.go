package routes

import (
	"github.com/controllers"
	"github.com/gin-gonic/gin"
)

func ApplicationV1Router(router *gin.Engine) {

	v1 := router.Group("/v1")
	{
		v1.POST("/user", controllers.CreateUser)
		v1.POST("/login2", controllers.LoginUser)
	}

}
