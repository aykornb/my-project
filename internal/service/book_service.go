package service

import "my-project/internal/domain"

type BookService struct {
	repo domain.BookRepository // Interface injection
}

func NewBookService(r domain.BookRepository) *BookService {
	return &BookService{repo: r}
}

func (s *BookService) FetchBook(id string) (*domain.Book, error) {
	// Business logic goes here (validation, caching, etc.)
	return s.repo.GetByID(id)
}

func (s *BookService) CreateBook(book *domain.Book) error {
	return s.repo.Save(book)
}

func (s *BookService) FetchAllBooks() ([]*domain.Book, error) {
	// Business logic goes here (validation, caching, etc.)
	return s.repo.GetAll()
}
