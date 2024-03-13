package routes

import (
	"bookmyshow/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, access-control-allow-origin")

		// If the request method is OPTIONS, handle it and return 200 status code
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// Continue to the next middleware or route handler
		c.Next()
	})

	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/signup", handlers.Signup)
	r.POST("/auth/refresh", handlers.RefreshToken)
	r.POST("/organizer/show", handlers.CreateShow)
	r.POST("/organizer/nfts", handlers.AddOrganizerNfts)
	r.POST("/shows", handlers.GetAllShows)
	r.GET("/test", handlers.Test)
}
