package main

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

func connect() {
	opts := &pg.Options{
		User:     "postgres",
		Password: "",
		Addr:     "localhost:5432",
	}
	db := pg.Connect(opts)

  if db == nil{
    log.Println("Failed to connect to the database")
    os.Exit()
  }
  log.Printf("Connection to the database is succesful")
}
