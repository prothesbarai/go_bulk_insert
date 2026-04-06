package database

import (
	"context"
	"database/sql"
	"go_bulk_insert/logger"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB(){
	dsn := os.Getenv("DB_DSN")
	db, err := sql.Open("mysql",dsn)
	if (err != nil) {
		logger.AppLogger.Error.Println("Failed Database Connect : ",err)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil{
		logger.AppLogger.Error.Println("Database Unrechable : ",err)
		os.Exit(1)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5*time.Minute)
	DB=db
	logger.AppLogger.Info.Println("Successfully Connected Database : ",db)
}