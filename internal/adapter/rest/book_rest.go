package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kritpi/go-clearn-arch/entities"
	"github.com/kritpi/go-clearn-arch/exceptions"
	"github.com/kritpi/go-clearn-arch/usecases"
)

type bookRestHandler struct {
	bookUseCase usecases.BookUseCases
}

func NewBookRestHandler(bookUseCase usecases.BookUseCases) bookRestHandler {
	return bookRestHandler{
		bookUseCase: bookUseCase,
	}
}

func (b *bookRestHandler) GetAllBooks(c *fiber.Ctx) error {
	books := b.bookUseCase.GetAllBooks()

	return c.Status(fiber.StatusOK).JSON(books)
}

func (b *bookRestHandler) CreateBook(c *fiber.Ctx) error {
	var book entities.Book

	err := c.BodyParser(&book)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse book",
		})
	}

	createdBook, err := b.bookUseCase.CreateBook(book)
	if err != nil {
		switch err {
		case exceptions.ErrBookIdDup:
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error": "this id is exist",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error",
			})
		}
	}

	return c.Status(fiber.StatusCreated).JSON(createdBook)
}