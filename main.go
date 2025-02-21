package main

import (
	"context"
	_ "database/sql"
	"elibrary/internal/delivery/bookhandlers"
	"elibrary/internal/delivery/userhandlers"
	"elibrary/internal/repository/bookrepository"
	"elibrary/internal/repository/userrepository"
	"elibrary/internal/usecase/bookusecase"
	"elibrary/internal/usecase/userusecase"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"

	"elibrary/internal/delivery"
)

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://atajan:postgres@localhost:5432/bookdb?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Fatal("Failed to close connection:", err)
		}
	}(conn, context.Background())

	bookRepo := bookrepository.NewBookRepository(conn)
	userRepo := userrepository.NewUserRepository(conn)

	bookUsecase := bookusecase.NewBookUsecase(bookRepo)
	userUsecase := userusecase.NewUserUsecase(userRepo)

	bookHandler := bookhandlers.NewBookHandler(bookUsecase)
	userHandler := userhandlers.NewUserHandler(userUsecase)

	router := mux.NewRouter()

	router.HandleFunc("/register", userHandler.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", userHandler.LoginHandler).Methods("POST")

	router.Handle("/books", delivery.AuthMiddleware("superadmin", "admin")(http.HandlerFunc(bookHandler.GetBooksHandler))).Methods("GET")
	router.Handle("/books/{id}", delivery.AuthMiddleware("superadmin", "admin")(http.HandlerFunc(bookHandler.DownloadBookHandler))).Methods("GET")

	router.Handle("/books", delivery.AuthMiddleware("superadmin")(http.HandlerFunc(bookHandler.UploadBookHandler))).Methods("POST")

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
