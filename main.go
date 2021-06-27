package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Book Struct = Model
type Book struct {
	ID			string		`json:"id"`
	Isbn		string		`json:"isbn"`
	Title		string		`json:"title"`
	Author		*Author		`json:"author"`
}

//Author struct
type Author struct {
	Firstname 	string 		`json:"firstname"`
	Lastname 	string 		`json:"lastname"`
}



//init books var as slice book struct
var books []Book

//Get books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	//Loop through books and find with ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

//Create book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books,book) 
	json.NewEncoder(w).Encode(book)
}

//Update book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books,book) 
			json.NewEncoder(w).Encode(book)
			return
		}
	}

	
}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func main() {
	//Init router
	r := mux.NewRouter()


	//Mock data
	books = append(books, Book{ID: "1", Isbn: "14523534",Title: "Book One", Author: &Author{Firstname: "john", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "145343534",Title: "Book Two", Author: &Author{Firstname: "jon", Lastname: "Doewwww"}})

	//Route handlers
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Sercver is starting...")
	log.Fatal(http.ListenAndServe(":8000",r))

	
}
