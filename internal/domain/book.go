package domain

type Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// The Service layer will use this.
type BookRepository interface {
	GetByID(id string) (*Book, error)
	Save(book *Book) error
	//     Update(book *Book) error
	GetAll() ([]*Book, error)
	//    CreateBook(book *Book) error
	//     Delete(id string) error
}
