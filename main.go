package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/go-clearn-arch/internal/adapter/memory"
	"github.com/kritpi/go-clearn-arch/internal/adapter/rest"
	"github.com/kritpi/go-clearn-arch/usecases"
)

func main() {
	bookRepo := memory.NewBookMemory()
	bookService := usecases.NewBookService(bookRepo)
	bookHandler := rest.NewBookRestHandler(bookService)

	app := fiber.New()

	app.Get("/books", bookHandler.GetAllBooks)
	app.Post("/books", bookHandler.CreateBook)

	app.Listen(":3000")
}