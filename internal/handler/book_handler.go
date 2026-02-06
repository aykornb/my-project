package handler

import (
	"encoding/json"
	"net/http"

	"my-project/internal/domain"
	"my-project/internal/service"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(s *service.BookService) *BookHandler {
	return &BookHandler{service: s}
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	book, err := h.service.FetchBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var b domain.Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateBook(&b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.FetchAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
