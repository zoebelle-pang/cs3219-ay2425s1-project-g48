package routes

import (
	controller "backend/controllers"

	"github.com/gin-gonic/gin"
)

// UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/v1/signup", controller.SignUp())
	incomingRoutes.POST("/v1/login", controller.Login())
	incomingRoutes.POST("/v1/refresh", controller.RefreshToken)
}
