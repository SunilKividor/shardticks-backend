package routes

import (
	"bookmyshow/handlers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/signup", handlers.Signup)
	r.POST("/auth/refresh", handlers.RefreshToken)
	r.GET("/test", handlers.Test)
}
