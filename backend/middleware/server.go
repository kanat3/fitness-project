package middleware

import (
	"fitness-project/backend/internal/storage"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getParam(c *gin.Context, paramName string) string {
	return c.Params.ByName(paramName)
}

func Account(c *gin.Context) {

	const op = "server.account"

	req := getParam(c, "Id")
	id, err := strconv.Atoi(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "params converting error", "from": op})
		return
	}
	// set storage
	var db *storage.Storage = storage.StoragePSQL

	existingUser, err := db.IsUserByID(id)

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

	req := getParam(c, "Id")
	id, err := strconv.Atoi(req)

	if err != nil {
		c.JSON(400, gin.H{"error": "params converting error", "from": op})
		return
	}

	// set storage
	var db *storage.Storage = storage.StoragePSQL
	existingUser, err := db.IsUserByID(id)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist", "from": op})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	var bid storage.Bid

	if err := c.ShouldBindJSON(&bid); err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	t := time.Now()
	t.Format("2006-01-02 15:04:05")
	bid.Created = t

	id_bid, err := db.SetStructBid(bid)

	if id_bid == 0 {
		c.JSON(400, gin.H{"error": "bid does not exist", "from": op})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	var ubid storage.BidByUser
	ubid.RefBid = id_bid
	ubid.RefUsers = existingUser.ID

	err = db.SetStructBidByUser(ubid)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	c.JSON(200, gin.H{"success": "bid added", "from": op})
}

/* get last your bid */
func GetBid(c *gin.Context) {

	const op = "server.getbid"

	// set storage
	var db *storage.Storage = storage.StoragePSQL

	req := getParam(c, "Id")
	id, err := strconv.Atoi(req)
	if err != nil {
		c.JSON(400, gin.H{"error": "params converting error", "from": op})
		return
	}

	existingUser, err := db.IsUserByID(id)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist", "from": op})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	var bid storage.Bid

	bid, err = db.GetLastBidByUser(id)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "from": op})
		return
	}

	c.JSON(200, bid)

}
