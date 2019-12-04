package main

import (
	"log"
	"math/rand"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

/**
 * Author struct
 */
type Author struct {
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}

/**
 * Book struct (As a Model here)
 */
type Book struct {
	ID string `json: "id"`
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Author *Author `json: "author"`
}

/**
 * Globally defining Book as slice Book struct
 */
var books []Book

/**
 * Get all books
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request        [description]
 * @return Jsonresponse
 */
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

/**
 * Get single books
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request        [description]
 * @return Jsonresponse
 */
func getSingleBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) //Getting the params

	//Getting the correct book by ID
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

/**
 * Create a book
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request        [description]
 * @return Jsonresponse
 */
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(10)) //Mocking ID- Not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

/**
 * Update a book
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request        [description]
 * @return Jsonresponse
 */
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) //fetching params

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book

			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"] //random id
			books = append(books, book)

			json.NewEncoder(w).Encode(book)
			return
		}
	}

	json.NewEncoder(w).Encode(books)
}

/**
 * Delete a book
 * @param  {[type]} w http.ResponseWriter [description]
 * @param  {[type]} r *http.Request        [description]
 * @return Jsonresponse
 */
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) //retriving ID

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

/**
 * Initialising router
 * @return {void}
 */
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	/**
	 * Handling the endpoints
	 */
	router.HandleFunc("/api/v1/books"		, getAllBooks).Methods("GET")
	router.HandleFunc("/api/v1/books/{id}"	, getSingleBook).Methods("GET")
	router.HandleFunc("/api/v1/books"		, createBook).Methods("POST")
	router.HandleFunc("/api/v1/books/{id}"	, updateBook).Methods("PUT")
	router.HandleFunc("/api/v1/books/{id}"	, deleteBook).Methods("DELETE")
	
	/** Logging the mesage for serving at localhost  */
	log.Fatal(http.ListenAndServe(":8087", router))
}

func main() {
	books = append(books, Book{ID: "1", Isbn: "36656", Title: "Shahrukh Anwar", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "54551", Title: "Tony Stark", Author: &Author{Firstname: "Jone", Lastname: "Foster"}})

	handleRequests()
}