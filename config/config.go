package config

import (
	"database/sql"
	"fmt"
	"time"
)

func InitDatabase(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("error InitDatabase", err)
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		panic(err)
	}

	fmt.Println("Connect to the database3 succesfully")

	return db
}
