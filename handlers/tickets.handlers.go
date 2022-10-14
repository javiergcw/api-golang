package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/faztweb/golang-restapi-crud/data"
	"github.com/faztweb/golang-restapi-crud/models"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

// POST
func CreateTickets(w http.ResponseWriter, r *http.Request) {

	var newTicket models.Ticket
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Inserta un ticket valido")
	}

	json.Unmarshal(reqBody, &newTicket)
	newTicket.ID = len(data.TicketData) + 1
	data.TicketData = append(data.TicketData, newTicket)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newTicket)
	if err != nil {
		return
	}
}

// END

// GET
func GetTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data.TicketData)
	if err != nil {
		return
	}
}

func GetOneTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketId, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "ID Invalido")
		return
	}

	for _, ticket := range data.TicketData {
		if ticket.ID == ticketId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ticket)
		}
	}

}

// END

// Update
func UpdatedTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketId, err := strconv.Atoi(vars["id"])
	var updatedTicket models.Ticket

	if err != nil {
		fmt.Fprintf(w, "ID invalido")
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Ingrese un dato valido")
	}
	json.Unmarshal(reqBody, &updatedTicket)

	for i, t := range data.TicketData {
		if t.ID == ticketId {
			data.TicketData = append(data.TicketData[:i], data.TicketData[i+1:]...)
			updatedTicket.ID = ticketId
			data.TicketData = append(data.TicketData, updatedTicket)

			fmt.Fprintf(w, "El Ticket con el ID %v ha sido acutalizada con éxito", ticketId)
		}
	}
}

// END

// DELETE
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ticketID, err := strconv.Atoi(
		vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Inserta un ID Valido")
		return
	}

	for i, t := range data.TicketData {
		if t.ID == ticketID {
			data.TicketData = append(data.TicketData[:i], data.TicketData[i+1:]...)
			fmt.Fprintf(w, "El ticket con el ID %v ha sido removido con éxito", ticketID)
		}
	}
}

// END

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello world"))
}
