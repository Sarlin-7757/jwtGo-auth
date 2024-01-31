package routes

import (
	middleware "github.com/Sarlin-7757/jwtGo-auth/middleware"

	controller "github.com/Sarlin-7757/jwtGo-auth/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
