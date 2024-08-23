package routers

import (
	"grpc-gateway/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, controllers *controllers.UserController) {
	group := router.Group("/user")
	group.GET("/:id", controllers.GetUserController)
	group.POST("", controllers.CreateUserController)
	group.PUT("", controllers.UpdateUserController)
	group.DELETE("/:id", controllers.DeleteUserController)
}
