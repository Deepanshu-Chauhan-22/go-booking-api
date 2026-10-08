// package main

// import (
// 	"bookingapp/booking"
// 	collectdata "bookingapp/collectData"
// 	"bookingapp/models"
// 	sendemail "bookingapp/sendEmail"
// 	"bookingapp/validation"
// 	"fmt"
// 	"sync"
// )

// func greet() {
// 	fmt.Println("Welcome to our conference page")
// 	fmt.Println("Book your tickets here")
// }

// func main() {
// 	var wg sync.WaitGroup
// 	greet()
// 	ch := make(chan models.User)
// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		go sendemail.SendEmail(ch)
// 	}()

// 	for booking.RemainingTickets > 0 {
// 		name, email, tickets := collectdata.CollectData()

// 		err := validation.ValidateInputs(name, email, tickets)
// 		if err != nil {
// 			fmt.Println("Error occured - ", err)
// 		} else {
// 			user, err := booking.Booking(name, email, tickets)
// 			if err != nil {
// 				fmt.Println("Error occured - ", err)
// 			} else {
// 				ch <- user
// 			}
// 		}
// 	}
// 	close(ch)
// 	wg.Wait()
// }

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /health", healthHandler)

	fmt.Println("Server starting on http://localhost:8080")

	http.HandleFunc("POST /bookings", createBookingHandler)
	fmt.Println("Directing to Booking Handler")

	http.HandleFunc("GET /tickets", ticketsHandler)
	fmt.Println("Getting the remaining tickets -")

	http.HandleFunc("GET /bookings", getBookingsHandler)
	fmt.Println("Getting all the bookings :")

	http.HandleFunc("GET /bookings/{email}", bookingsEmailHandler)
	fmt.Println("Getting the booking details of this email - ")

	http.HandleFunc("PUT /bookings/{email}", updateTicketsHandler)
	fmt.Println("Updating the user tickets")

	http.HandleFunc("DELETE /bookings/{email}", cancelBookingHandler)
	fmt.Println("Cancelling booking for user ")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
