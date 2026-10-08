package sendemail

import (
	"bookingapp/models"
	"fmt"
	"time"
)

func SendEmail(ch chan models.User) {
	//defer models.Wg.Done()
	for user := range ch {
		time.Sleep(10 * time.Second)
		fmt.Println("########################")
		fmt.Printf("\nThank you %v ! This is the confirmation of your conference booking\n", user.Name)
		fmt.Printf("\nYou have booked %v tickets on the registered email %v\n ", user.TicketsBought, user.Email)
		fmt.Println("########################")
	}
}
