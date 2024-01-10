package middleware

import (
	"fitness-project/backend/internal/storage"

	"github.com/gin-gonic/gin"
)

func Account(c *gin.Context) {

	const op = "server.account"
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

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	c.JSON(200, existingUser)
}

func SetBid(c *gin.Context) {

	const op = "server.setbid"

	// set storage
	var db *storage.Storage = storage.StoragePSQL

	var bid storage.Bid

	if err := c.ShouldBindJSON(&bid); err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	err := db.SetStructBid(bid)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	c.JSON(200, gin.H{"success": "bid added", "from": op})
}

func GetBid(c *gin.Context) {

	const op = "server.getbid"

	// set storage
	var db *storage.Storage = storage.StoragePSQL
	var user storage.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	existingUser, err := db.IsUserByEmail(user.Email)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist", "from": op})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	var bid storage.Bid

	bid, err = db.GetLastBidByUser(user.ID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	c.JSON(200, bid)

}
