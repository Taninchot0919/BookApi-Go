package book

import (
	"github.com/gofiber/fiber"
)

func GetBooks(context *fiber.Ctx) {
	context.Send("All Books")
}

func GetBook(context *fiber.Ctx) {
	context.Send("A Singular Book")
}

func NewBook(context *fiber.Ctx) {
	context.Send("Add a new book")
}

func DeleteBook(context *fiber.Ctx) {
	context.Send("Delete a book")
}
