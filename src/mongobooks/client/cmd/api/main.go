package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"mongobooks_client/internal/books"
	"mongobooks_client/pkg/db"

	"github.com/gorilla/mux"
)

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("ENV \"%s\" not found, default value \"%s\" used", key, defaultVal)
	return defaultVal
}

func main() {
	db_user := getEnv("DB_USER", "admin")
	db_pass := getEnv("DB_PASS", "password")
	db_host := getEnv("DB_HOST", "localhost")
	db_port := getEnv("DB_PORT", "27017")

	db_con_str := fmt.Sprintf("mongodb://%s:%s@%s:%s", db_user, db_pass, db_host, db_port)
	database := db.Connect(db_con_str)

	repo := books.NewRepository(database)
	handler := books.NewHandler(repo)

	r := mux.NewRouter()
	r.HandleFunc("/books", handler.GetBooks).Methods("GET")
	r.HandleFunc("/books", handler.CreateBook).Methods("POST")

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
