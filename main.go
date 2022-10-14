package main

import (
	"fmt"
	"github.com/faztweb/golang-restapi-crud/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// Index Routes
	router.HandleFunc("/", handlers.IndexRoute)

	// Tickets Routes
	router.HandleFunc("/tickets", handlers.GetTickets).Methods("GET")
	router.HandleFunc("/tickets/{id}", handlers.GetOneTicket).Methods("GET")
	router.HandleFunc("/tickets", handlers.CreateTickets).Methods("POST")
	router.HandleFunc("/tickets/{id}", handlers.DeleteTicket).Methods("DELETE")
	router.HandleFunc("/tickets/{id}", handlers.UpdatedTicket).Methods("PUT")

	fmt.Println("Servidor iniciado en el puerto", 3000)
	log.Fatal(http.ListenAndServe(":3000", router))

}
