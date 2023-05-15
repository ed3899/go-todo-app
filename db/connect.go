package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/edca3899/go-todo-mysql/config"
	"github.com/go-sql-driver/mysql"
)

var (
	db_address  string
	db_username string
	db_password string
	db_name     string
)

func init() {
	db_address = config.Config("DB_ADDRESS")
	db_username = config.Config("DB_USER")
	db_password = config.Config("DB_PASSWORD")
	db_name = config.Config("DB_NAME")
}

var DB *sql.DB

func ConnectDB() {
	var err error

	cfg := mysql.Config{
		Addr:   db_address,
		User:   db_username,
		Passwd: db_password,
		DBName: db_name,
		Net:    "tcp",
		AllowNativePasswords: true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("Couldn't connect to db because of: %v", err)
	}

	if err := DB.PingContext(ctx); err != nil {
		log.Fatalf("Couldn't ping database because of: %v", err)
	} else {
		log.Print("Successfully pinged db")
	}
}
