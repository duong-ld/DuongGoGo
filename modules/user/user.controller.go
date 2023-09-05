package user

import (
	"github.com/gin-gonic/gin"
)

var service = new(Service)

func Routes(routerGroup *gin.RouterGroup, middleware ...gin.HandlerFunc) {
	userRouterGroup := routerGroup.Group("user")
	userRouterGroup.Use(middleware...)
	userRouterGroup.GET("/", service.GetAll)
	userRouterGroup.POST("/", service.CreateUser)
	userRouterGroup.GET("/:id", service.GetUser)
	userRouterGroup.PUT("/:id", service.UpdateUser)
}
