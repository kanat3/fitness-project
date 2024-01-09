package api

import (
	"fitness-project/backend/middleware"

	"github.com/gin-gonic/gin"
)

func InitHandlers(r *gin.Engine) {

	r.POST("/login", middleware.Login)
	r.POST("/signup", middleware.Signup)
	r.POST("/reset-password", middleware.ResetPassword)
	r.POST("/reset-email", middleware.ResetEmail)

	r.GET("/logout", middleware.Logout)
	r.GET("/status", Status)
	r.GET("/login", Login)
	r.GET("/signup", Signup)

	/* add front */
	r.LoadHTMLGlob("frontend/*.html")
	r.Static("frontend", "./frontend")
	//	r.Static("frontend/assets/css", "./frontend/assets/css")
	//	r.Static("frontend/assets/js", "./frontend/assets/js")

	r.GET("/about", About)
	r.GET("/trainers", Trainers)
	r.GET("/contacts", Contacts)
	r.GET("/account", Account)
	r.GET("/home", Home)
	r.POST("/test", middleware.IsAuthorized(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Test request created"})
	})

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
	c.HTML(200, "reg.html", nil)
}

func About(c *gin.Context) {
	c.HTML(200, "about_us.html", nil)
}

func Contacts(c *gin.Context) {
	c.HTML(200, "contacts.html", nil)
}

func Account(c *gin.Context) {
	c.HTML(200, "personal_account.html", nil)
}

func Trainers(c *gin.Context) {
	c.HTML(200, "trainers.html", nil)
}

func Home(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
