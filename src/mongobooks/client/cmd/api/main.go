package main

import (
	"log"
	"net/http"

	"mongobooks_client/internal/books"
	"mongobooks_client/pkg/db"

	"github.com/gorilla/mux"
)

func main() {
	database := db.Connect("mongodb://admin:password@localhost:27017")

	repo := books.NewRepository(database)
	handler := books.NewHandler(repo)

	r := mux.NewRouter()
	r.HandleFunc("/books", handler.GetBooks).Methods("GET")
	r.HandleFunc("/books", handler.CreateBook).Methods("POST")

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
