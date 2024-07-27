package repositories

import "github.com/kritpi/go-clearn-arch/entities"

type BookRepository interface {
	GetAll() []entities.Book
	Create(book entities.Book) (*entities.Book, error)
}
