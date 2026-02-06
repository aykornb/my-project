package repository

import (
	"errors"
	"sync"

	"my-project/internal/domain" // Import our "Contract"
)

// InMemoryBookRepo is our "Class" equivalent.
// It's a struct that holds the state (the map and the lock).
type InMemoryBookRepo struct {
	mu    sync.RWMutex
	books map[string]*domain.Book
}

// NewInMemoryRepo is our "Constructor."
// In Go, it's idiomatic to return a pointer to the struct.
func NewInMemoryRepo() *InMemoryBookRepo {
	return &InMemoryBookRepo{
		books: make(map[string]*domain.Book),
	}
}

// GetByID implements the domain.BookRepository interface.
// Notice the (r *InMemoryBookRepo) receiver - this is how you "attach" methods.
func (r *InMemoryBookRepo) GetByID(id string) (*domain.Book, error) {
	r.mu.RLock() // Read Lock (multiple readers allowed)
	defer r.mu.RUnlock()

	book, ok := r.books[id]
	if !ok {
		return nil, errors.New("book not found")
	}

	return book, nil
}

// Save implements the domain.BookRepository interface.
func (r *InMemoryBookRepo) Save(book *domain.Book) error {
	r.mu.Lock() // Write Lock (exclusive access)
	defer r.mu.Unlock()

	r.books[book.ID] = book
	return nil
}

func (r *InMemoryBookRepo) GetAll() ([]*domain.Book, error) {
	r.mu.RLock() // Read Lock (multiple readers allowed)
	defer r.mu.RUnlock()
	books := make([]*domain.Book, 0, len(r.books))
	for _, book := range r.books {
		books = append(books, book)
	}

	return books, nil
}
