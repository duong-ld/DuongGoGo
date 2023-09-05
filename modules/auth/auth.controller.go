package auth

import (
	"duongGoGo/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(routerGroup *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	authRouter := routerGroup.Use(middlewares...)
	authRouter.POST("/sign-up", authService.SignUp)
	authRouter.POST("/sign-in", authService.SignIn)
	authRouter.POST("/logout", middleware.JWTMiddleware(), middleware.CSRFMiddleware(), authService.Logout)
}
