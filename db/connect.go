package db

import (
	"database/sql"
	"log"


	"github.com/edca3899/go-todo-mysql/config"
	"github.com/go-sql-driver/mysql"
)

var (
	db_username string
	db_password string
	db_address  string
	db_name     string
)

func init() {
	db_username = config.Config("DB_USER")
	db_password = config.Config("DB_PASSWORD")
	db_address = config.Config("DB_ADDRESS")
	db_name = config.Config("DB_NAME")
}

var DB *sql.DB

func ConnectDB() {
	var err error

	cfg := mysql.Config{
		User:   db_username,
		Passwd: db_password,
		Net:    "tcp",
		Addr:   db_address,
		DBName: db_name,
	}

	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Couldn't connect to db because of: %v", err)
	}
}
