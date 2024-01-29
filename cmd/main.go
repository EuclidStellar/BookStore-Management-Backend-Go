package main

import (
	"bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RigesterBookStoreRooutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))

}