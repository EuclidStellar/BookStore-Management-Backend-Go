package routes

import (
	
	"bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RigesterBookStoreRooutes = func(router *mux.Router) {
	router.HandleFunc("/book/" , controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", controllers.DeleteBook).Methods("DELETE")

	


}

