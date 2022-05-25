package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Ticket struct {
	ID    string `json:"id"`
	Isbn  string `json:"Isbn"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var tickets []Ticket

func main() {
	tickets = append(tickets, Ticket{ID: "1", Isbn: "9009", Name: "Astana", Price: 15000})
	tickets = append(tickets, Ticket{ID: "2", Isbn: "4508", Name: "Moscow", Price: 22000})
	request := mux.NewRouter()
	request.HandleFunc("/tickets", getTickets).Methods("GET")
	request.HandleFunc("/tickets/{id}", getTicket).Methods("GET")
	request.HandleFunc("/tickets", createTicket).Methods("POST")
	request.HandleFunc("/tickets/{id}", updateTicket).Methods("PUT")
	request.HandleFunc("/tickets/{id}", deleteTicket).Methods("DELETE")

	fmt.Printf("Server started at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", request))
}

func getTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

func getTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range tickets {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var ticket Ticket
	_ = json.NewDecoder(r.Body).Decode(&ticket)
	ticket.ID = strconv.Itoa(rand.Intn(10000))
	tickets = append(tickets, ticket)
	json.NewEncoder(w).Encode(ticket)
}

func updateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range tickets {
		if item.ID == params["id"] {
			tickets = append(tickets[:index], tickets[index+1:]...)
			var ticket Ticket
			_ = json.NewDecoder(r.Body).Decode(&ticket)
			ticket.ID = params["id"]
			tickets = append(tickets, ticket)
			json.NewEncoder(w).Encode(ticket)
			return
		}

	}

}

func deleteTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range tickets {
		if item.ID == params["id"] {
			tickets = append(tickets[:index], tickets[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(tickets)
}
