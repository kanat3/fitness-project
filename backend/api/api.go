package api

import (
	"fitness-project/backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitHandlers(r *gin.Engine) {
	r.POST("/login", middleware.Login)
	r.POST("/signup", middleware.Signup)
	r.GET("/logout", middleware.Logout)
	r.GET("/status", Status)
}

// status
func Status(c *gin.Context) {
	const op = "api.status"
	c.JSON(200, gin.H{"status": "ok", "from": op})
}
