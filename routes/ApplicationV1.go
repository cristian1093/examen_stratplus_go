package routes

import (
	"github.com/controllers"
	"github.com/gin-gonic/gin"
)

func ApplicationV1Router(router *gin.Engine) {

	v1 := router.Group("/v1")

	{

		v1.POST("/login", controllers.LoginUser)

		v1.POST("/user", controllers.CreateUser)
		v1.PUT("/user/:id", controllers.TokenMiddleware(), controllers.UpdateUser)
		v1.GET("/user", controllers.TokenMiddleware(), controllers.ShowUser)
		v1.DELETE("/user/:id", controllers.TokenMiddleware(), controllers.DeleteUser)

		v1.POST("/create-order", controllers.TokenMiddleware(), controllers.CreateOrder)
		v1.GET("/orders/:id", controllers.TokenMiddleware(), controllers.GetUserOrders)
		v1.GET("/order-book", controllers.TokenMiddleware(), controllers.GetOrderBook)
		v1.PUT("/update-order/:id", controllers.TokenMiddleware(), controllers.UpdateOrder)

		v1.POST("/bond", controllers.TokenMiddleware(), controllers.CreateBond)
		v1.GET("/bond", controllers.TokenMiddleware(), controllers.ShowBond)
		v1.PUT("/bond/:id", controllers.TokenMiddleware(), controllers.UpdateBond)
		v1.DELETE("/bond/:id", controllers.TokenMiddleware(), controllers.DeleteBond)
	}

}
