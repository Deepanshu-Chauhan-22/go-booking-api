package main

import (
	"bookingapp/booking"
	"bookingapp/models"
	"bookingapp/validation"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Tickets struct {
	Ticket int `json:"remainingTicket"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HealthResponse{
		Status:  "ok",
		Message: "Booking API is running",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Failed to write health response:", err)
	}
}

func createBookingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error while decoding JSON into struct")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// now we have the decoded data ... lets convert it into var
	name := user.Name
	email := user.Email
	tickets := user.TicketsBought
	err = validation.ValidateInputs(name, email, tickets)
	if err != nil {
		fmt.Println("Invalid Input! , ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// now we will call the booking function because the inputs are valid
	user, err = booking.Booking(name, email, tickets)
	if err != nil {
		fmt.Println("Error while booking ! ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

func ticketsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tickets := booking.GetTickets()

	err := json.NewEncoder(w).Encode(&tickets)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error while Encoding - ", err)
		return
	}
}

func getBookingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var mem []models.User
	mem = booking.GetBookings()
	err := json.NewEncoder(w).Encode(mem)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error while Encoding - ", err)
		return
	}

}

func bookingsEmailHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//userEmailData := booking.UserDatabyEmail()

	usrEmail := r.PathValue("email")

	//user, ok := userEmailData[usrEmail]

	user, er := booking.UserDatabyEmail(usrEmail)

	if er != nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(user)

	if err != nil {
		fmt.Println("Error occured while Encoding - ", err)
		return
	}
}

func updateTicketsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	usrEmail := r.PathValue("email")
	var newTickets int
	err := json.NewDecoder(r.Body).Decode(&newTickets)

	if err != nil {
		fmt.Println("Error while decoding the request ", err)
		return
	}

	usr, err := booking.UpdateTickets(usrEmail, newTickets)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(usr); err != nil {
		log.Println("Error encoding updated booking:", err)
	}
}

func cancelBookingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	usrEmail := r.PathValue("email")

	usr, err := booking.CancelBooking(usrEmail)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(usr)
	if err != nil {
		log.Println("Error encoding deleting booking:", err)
	}

}
