package book

import (
	"github.com/elliotforbes/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(context *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	context.JSON(books)
}

func GetBook(context *fiber.Ctx) {
	id := context.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	context.JSON(book)
}

func NewBook(context *fiber.Ctx) {
	db := database.DBConn
	book := new(Book)
	if err := context.BodyParser(book); err != nil {
		context.Status(503).Send(err)
		return
	}
	db.Create(&book)
	context.JSON(book)
}

func DeleteBook(context *fiber.Ctx) {
	id := context.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		context.Status(500).Send("No book from the given ID")
		return
	}
	db.Delete(&book)
	context.Send("Delete Successfully")
}
