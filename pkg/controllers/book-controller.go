package controllers

import (
	models "bookstore/pkg/models"
	"bookstore/pkg/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book
func GetBook(w http.ResponseWriter, r *http.Request) {
	newbooks:= models.GetAllBooks()
	res, _ := json.Marshal(newbooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)


}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID , err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		panic(err)
	}	
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.Parsebody(r, CreateBook)
	bookDetails := CreateBook.CreateBook()
 	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID , err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		panic(err)
	}	
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatebook = &models.Book{}
	utils.Parsebody(r, updatebook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID , err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		panic(err)
	}	
	bookDetails, db := models.GetBookById(ID)
	if updatebook.Name != "" {
		bookDetails.Name = updatebook.Name
	}
	if updatebook.Author != "" {
		bookDetails.Author = updatebook.Author
	}
	if updatebook.Publication != "" {
		bookDetails.Publication = updatebook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}