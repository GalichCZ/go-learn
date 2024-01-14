package router

import (
	"crud-db/pkg/controller"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book", controller.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{id}", controller.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{id}", controller.UpdateBook).Methods("PUT")
}
