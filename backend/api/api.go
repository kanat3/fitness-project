package api

import (
	"fitness-project/backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitHandlers(r *gin.Engine) {
	// test html
	//r.LoadHTMLFiles("static/login.html")
	//r.Static("static/css", "./static/css")

	r.POST("/login", middleware.Login)
	r.POST("/signup", middleware.Signup)
	r.POST("/reset-password", middleware.ResetPassword)
	r.POST("/reset-email", middleware.ResetEmail)

	r.GET("/logout", middleware.Logout)
	r.GET("/status", Status)
	r.GET("/login", Login)
	r.GET("/signup", Signup)
}

func Status(c *gin.Context) {
	const op = "api.status"
	c.JSON(200, gin.H{"status": "ok", "from": op})
}

/* TODO */
func Login(c *gin.Context) {
	c.JSON(200, gin.H{"status": "test GET"})
	//c.HTML(200, "login.html", nil)
}

/* TODO */
func Signup(c *gin.Context) {
	c.JSON(200, gin.H{"status": "test GET"})
	//c.HTML(200, "signup.html", nil)
}
