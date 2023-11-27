package middleware

import (
	"fitness-project/backend/internal/storage"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseToken(tokenString string) (claims *Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}

func Login(c *gin.Context) {

	const op = "auth.login"
	fmt.Print("login")

	var user storage.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	// set storage
	var db *storage.Storage = storage.StoragePSQL

	existingUser, err := db.IsUserByEmail(user.Email)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist", "from": op})
		return
	}

	errHash := CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "invalid password", "from": op})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token", "from": op})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged in", "from": op})
}

func Logout(c *gin.Context) {

	const op = "auth.logout"

	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out", "from": op})
}

/* TODO: add check json, check password lengh */

func Signup(c *gin.Context) {

	const op = "auth.signup"

	var user storage.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	// set storage
	var db *storage.Storage = storage.StoragePSQL
	var existingUser storage.User

	existingUser, err := db.IsUserByEmail(user.Email)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "user already exists", "from": op})
		return
	}

	if user.Email == "" {
		c.JSON(400, gin.H{"error": "email isn't set", "from": op})
		return
	}

	var errHash error
	user.Password, errHash = GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash", "from": op})
		return
	}

	// set time
	t := time.Now()
	t.Format("2006-01-02 15:04:05")
	user.Created = t

	err = db.SetStructUsers(user)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	c.JSON(200, gin.H{"success": "user created", "from": op})
}

func IsAuthorized() gin.HandlerFunc {

	const op = "auth.is-authorized"
	return func(c *gin.Context) {
		cookie, err := c.Cookie("token")

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized", "from": op})
			c.Abort()
			return
		}

		_, err = ParseToken(cookie)

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized", "from": op})
			c.Abort()
			return
		}

		c.Next()
	}
}

/* TODO: add email notifications */

func ResetPassword(c *gin.Context) {

	const op = "auth.reset-password"
	var user storage.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//set storage
	var existingUser storage.User
	var db *storage.Storage = storage.StoragePSQL

	existingUser, err := db.IsUserByEmail(user.Email)

	if existingUser.ID == 0 || err != nil {
		c.JSON(400, gin.H{"error": "user does not exist", "from": op})
		return
	}

	var errHash error
	user.Password, errHash = GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	err = db.UsersUpdatePassword(existingUser.ID, user.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	c.JSON(200, gin.H{"success": "password updated"})
}

/* TODO: add email notifications */

func ResetEmail(c *gin.Context) {

	const op = "auth.reset-email"
	var user storage.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//set storage
	var existingUser storage.User
	var db *storage.Storage = storage.StoragePSQL

	existingUser, err := db.IsUserByPhone(user.Phone)

	if existingUser.ID == 0 || err != nil {
		c.JSON(400, gin.H{"error": "user does not exist", "from": op})
		return
	}

	if user.Email == "" {
		c.JSON(400, gin.H{"error": "email isn't set", "from": op})
		return
	}

	if user.Email == existingUser.Email {
		c.JSON(400, gin.H{"error": "new email is using. nothing to update", "from": op})
		return
	}

	err = db.UsersUpdateEmail(existingUser.ID, user.Email)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	c.JSON(200, gin.H{"success": "email updated"})
}
