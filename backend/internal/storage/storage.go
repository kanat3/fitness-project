package storage

import (
	"database/sql"
	"fmt"

	"fitness-project/backend/internal/config"

	_ "github.com/lib/pq"
)

const ()

type Storage struct {
	db         *sql.DB
	dbSettings config.DatabaseConfig
}

func (s *Storage) ConnectToDB() (err error) {
	const op = "storage.New"

	s.SetStorageSettings()

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		s.dbSettings.Host, s.dbSettings.Port, s.dbSettings.User, s.dbSettings.Password,
		s.dbSettings.DBname,
	)

	db, err := sql.Open("postgres", connStr)
	s.db = db

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(error.Error(err))
		return err
	}

	return nil
}

func (s *Storage) Close() {
	s.Close()
}

func (s *Storage) PingDB() error {

	s.SetStorageSettings()

	ping := "host=" + s.dbSettings.Host + " port=" + s.dbSettings.Port + " user=" + s.dbSettings.User +
		" password=" + s.dbSettings.Password + " sslmode=disable"
	db, error := sql.Open("postgres", ping)

	if error != nil {
		fmt.Println(error.Error())
		return error
	}

	error = db.Ping()

	return nil
}

func (s *Storage) SetStorageSettings() {
	cfg := config.GetConfig()
	s.dbSettings = cfg.DatabaseConfig
}

func (s *Storage) HealCheck() error {
	_, error := s.db.Query("select * from users;")

	if error != nil {
		fmt.Println(error.Error())
		return error
	}

	return nil
}
