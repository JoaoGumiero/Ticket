package api

import "net/http"

func HandleServer() *http.ServeMux{
	mux := http.NewServeMux()
	mux.HandleFunc("/", FirtPageHandler)
	mux.HandleFunc("/NewTicket", NewTicketPageHandler)
	mux.HandleFunc("GET /Tickets", GetAllTicketsHandler)
	mux.HandleFunc("GET /Ticket/{id}", GetTicketById) // Later can be substituted by a filter
	mux.HandleFunc("DELETE /Tikcet/{id}", DeleteTicketById)
	mux.HandleFunc("POST /Ticket", InsertTicketHandler)
	mux.HandleFunc("PATCH /Ticket/{id}", EditProductById)
	return mux
}