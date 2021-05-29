package main

import (
	"fmt"

	"github.com/elliotforbes/go-fiber-tutorial/book"
	"github.com/elliotforbes/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection Success!")
}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	app.Get("/", helloWorld)

	setupRoutes(app)

	app.Listen(3000)
}

func helloWorld(context *fiber.Ctx) {
	context.Send("Hello World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/book", book.GetBooks)
	app.Get("/api/book/:id", book.GetBook)
	app.Post("/api/book", book.NewBook)
	app.Delete("/api/book/:id", book.DeleteBook)
}
