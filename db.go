package main

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func Connect() {
	config := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "forum",
	}
	database, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected to database")
	db = database
}
