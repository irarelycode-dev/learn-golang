package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/irarelycode/golang-jwt/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	// incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUser)
	incomingRoutes.GET("/users/:user_id", controllers.GetUsers)
}
