package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

const (
	host     = ""
	port     = 
	user     = ""
	password = ""
	dbname   = ""
)

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos[] string
	rows, err := db.Query("SELECT * FROM todos")
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
		c.JSON("Ann Error Occured")
	}

	for rows.Next(){
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return c.Render("index", fiber.Map{
		"Todos":todos,
	})
}

func postHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func putHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//connecting to the DB
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := fiber.New()
	app.Get("/", indexHandler)
	app.Post("/", postHandler)
	app.Put("/update", putHandler)
	app.Delete("/", deleteHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
