package api

import (
	"encoding/json"
	. "github.com/floriank/saltmine/datastore"
	"net/http"
)

func (s *SaltmineAPI) TicketsList() func(http.ResponseWriter, *http.Request) {
	tickets, _ := s.tickets.List()

	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(tickets)
	}
}

func (s *SaltmineAPI) TicketGet() func(http.ResponseWriter, *http.Request) {
	tickets, _ := s.tickets.Get(1)
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(tickets)
	}
}

func (s *SaltmineAPI) TicketUpdate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&Ticket{})
	}
}

func (s *SaltmineAPI) TicketDelete() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&Ticket{})
	}
}

func (s *SaltmineAPI) TicketCreate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&Ticket{})
	}
}
