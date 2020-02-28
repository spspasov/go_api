package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Book struct {
	Id     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// A Slice of Books
var books []Book

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data
	books = append(books,
		Book{
			Id:    "1",
			Isbn:  "49034950",
			Title: "It",
			Author: &Author{
				FirstName: "Stephen",
				LastName:  "King",
			},
		}, Book{
			Id:    "2",
			Isbn:  "54903685",
			Title: "Jurassic Park",
			Author: &Author{
				FirstName: "Michael",
				LastName:  "Crichton",
			},
		}, Book{
			Id:    "3",
			Isbn:  "93851058",
			Title: "Ready Player One",
			Author: &Author{
				FirstName: "Earnest",
				LastName:  "Cline",
			},
		})

	// Route Handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	// TODO
	//r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	//r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

// GET /api/books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GET /api/book/1
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// POST /api/books
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(64))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
