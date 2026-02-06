package main

import (
	"log"
	"net/http"
	// These are your internal packages.
	// The prefix "my-project" must match your go.mod file.

	"my-project/internal/handler"
	"my-project/internal/repository"
	"my-project/internal/service"
)

func main() {
	// 1. Initialize Adapters (The "Data" layer)
	// Usually, this is where you'd open a DB connection.
	repo := repository.NewInMemoryRepo()

	// 2. Initialize Services (The "Logic" layer)
	// We inject the repo into the service.
	bookSvc := service.NewBookService(repo)

	// 3. Initialize Handlers (The "Transport" layer)
	// We inject the service into the handler.
	bookHandler := handler.NewBookHandler(bookSvc)

	// 4. Setup Router
	mux := http.NewServeMux()

	// Routing logic
	mux.HandleFunc("GET /books", bookHandler.GetAllBooks)
	mux.HandleFunc("GET /booksAll2", bookHandler.GetAllBooks)
	mux.HandleFunc("GET /books/{id}", bookHandler.GetBook)
	mux.HandleFunc("POST /books", bookHandler.CreateBook)

	// 5. Wrap router with Middleware (Interceptors)
	// This is the "Onion" pattern.
	loggingMux := handler.LoggingMiddleware(mux)

	// 6. Start Server
	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", loggingMux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
