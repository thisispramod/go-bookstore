package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thisispramod/go-bookstore/pkg/models"
	"github.com/thisispramod/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {

	// Database se saari books lao
	newBooks := models.GetAllBooks()

	// Go Struct -> JSON
	res, _ := json.Marshal(newBooks)

	// Response type JSON
	w.Header().Set("Content-Type", "application/json")

	// Status Code 200
	w.WriteHeader(http.StatusOK)

	// JSON client ko bhejo
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	// URL parameter nikalo
	vars := mux.Vars(r)

	// bookId string me milega
	bookId := vars["bookId"]

	// String -> Integer
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
		return
	}

	// Database se book nikalo
	bookDetails, _ := models.GetBookById(ID)

	// Struct -> JSON
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Create an empty Book struct
	updateBook := &models.Book{}

	// Parse JSON request body into updateBook
	utils.ParseBody(r, updateBook)

	// Get bookId from URL
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	// Convert string ID to int64
	ID, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid Book ID", http.StatusBadRequest)
		return
	}

	// Fetch existing book
	bookDetails, db := models.GetBookById(ID)

	// Update only provided fields
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	// Save updated book
	db.Save(&bookDetails)

	// Convert response to JSON
	res, err := json.Marshal(bookDetails)

	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	fmt.Println("Book updated successfully")
}
