package main

import (
	"encoding/json"
	"log"
	"net/http"

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

}

//Create book
func createBook(w http.ResponseWriter, r *http.Request) {

}

//Update book
func updateBook(w http.ResponseWriter, r *http.Request) {

}

//delete book
func deleteBook(w http.ResponseWriter, r *http.Request) {

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

	log.Fatal(http.ListenAndServe(":8000",r))

	
}
