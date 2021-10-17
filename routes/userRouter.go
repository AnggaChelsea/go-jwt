package routes

import (
	controller "go-jwt/controllers"
	"go-jwt/middleware"
	"github.com/gin-gonic/gin"
)
func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUser())
	incomingRoutes.GET("/user/:user_id", controller.GetUserId())
}