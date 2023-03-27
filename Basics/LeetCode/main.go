package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Create a new GORM database instance
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate the schema
	db.AutoMigrate(&Person{})

	// Define our routes
	app.Get("/people", func(c *fiber.Ctx) error {
		var people []Person
		db.Find(&people)
		return c.JSON(people)
	})

	app.Get("/people/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var person Person
		db.First(&person, id)
		return c.JSON(person)
	})

	app.Post("/people", func(c *fiber.Ctx) error {
		person := new(Person)
		if err := c.BodyParser(person); err != nil {
			return err
		}
		db.Create(&person)
		return c.JSON(person)
	})

	app.Put("/people/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var person Person
		db.First(&person, id)

		updatedPerson := new(Person)
		if err := c.BodyParser(updatedPerson); err != nil {
			return err
		}
		db.Model(&person).Updates(updatedPerson)
		return c.JSON(person)
	})

	app.Delete("/people/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var person Person
		db.Delete(&person, id)
		return nil
	})

	// Start the server
	app.Listen(":3000")
}
