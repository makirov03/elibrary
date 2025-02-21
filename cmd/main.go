package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"net/http"

	"elibrary/internal/delivery/bookhandlers"
	"elibrary/internal/delivery/userhandlers"
	"elibrary/internal/repository/bookrepository"
	"elibrary/internal/repository/userrepository"
	"elibrary/internal/usecase/bookusecase"
	"elibrary/internal/usecase/userusecase"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// db
	dbConn, err := pgx.Connect(context.Background(), "postgres://atajan:atajansan@localhost:5432/postgres")
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	bookRepo := bookrepository.NewBookRepository(dbConn)
	userRepo := userrepository.NewUserRepository(dbConn)

	// Use cases
	bookUC := bookusecase.NewBookUsecase(bookRepo)
	userUC := userusecase.NewUserUsecase(userRepo)

	// Handlers
	bookHandler := bookhandlers.NewBookHandler(bookUC)
	userHandler := userhandlers.NewUserHandler(userUC)

	// Routes (gorilla/mux)
	r.HandleFunc("/books", bookHandler.GetBooksHandler).Methods("GET")
	r.HandleFunc("/books/{id}", bookHandler.GetBooksHandler).Methods("GET")
	r.HandleFunc("/books", bookHandler.UploadBookHandler).Methods("POST")
	r.HandleFunc("/books/upload", bookHandler.UploadBookHandler).Methods("POST")

	// Routes
	r.HandleFunc("/users/register", userHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/users/login", userHandler.LoginHandler).Methods("POST")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
