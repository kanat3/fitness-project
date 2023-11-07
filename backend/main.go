package main

import (
	"fitness-project/backend/internal/config"
	"fitness-project/backend/internal/storage"
	"fmt"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init config
	cfg := config.GetConfig()

	_ = cfg

	// connect to database

	var db storage.Storage

	err := db.ConnectToDB()
	//defer db.Close()

	if err != nil {
		fmt.Println(error.Error(err))
		return
	}

	// healcheck db
	db.HealCheck()

	/* check setter

	var data storage.User
	data.FirstName = "anna2"
	data.SecondName = "a"
	data.LastName = "b"
	data.Phone = "89509190428"
	data.Email = "a.kuchebo@mail.ru"
	data.ProfileImg = "/"
	t := time.Now()
	t.Format("2006-01-02 15:04:05")
	data.Created = t

	fmt.Print(data.Created)

	err = db.SetStructUsers(data)
	if err != nil {
		fmt.Print(err)
	}
	*/

	// init router

	// run server
}
