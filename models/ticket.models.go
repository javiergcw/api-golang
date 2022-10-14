package models

type Ticket struct {
	ID           int    `json:"id"`
	Name         string `json:"Name"`
	Status       string `json:"Status""`
	CreationDate string `json:"CreationDate"`
	UpdateDate   string `json:"UpdateDate"`
}

type Tickets []Ticket
