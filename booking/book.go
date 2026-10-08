package booking

import (
	"bookingapp/models"
	"errors"
	"fmt"
	"sync"
)

var mu sync.Mutex

var RemainingTickets = 20000

var members []models.User

var UserData = make(map[string]models.User)

func Booking(name string, email string, tickets int) (models.User, error) {
	mu.Lock()
	defer mu.Unlock()

	_, ok := UserData[email]
	if ok {
		return models.User{}, errors.New(
			"email already has a booking; update the existing booking instead",
		)
	}
	var user models.User
	if tickets > RemainingTickets {
		err := errors.New("Sorry , the remaining tickets are less than the tickets you want")
		return user, err
	}
	//fmt.Println("Processing Booking ... ")
	RemainingTickets -= tickets
	user = models.User{
		Name:          name,
		Email:         email,
		TicketsBought: tickets,
	}
	members = append(members, user)
	UserData[user.Email] = user
	// fmt.Println("Thank you for booking the tickets")
	// fmt.Println("Remaining Tickets : ", RemainingTickets)
	// fmt.Println("Given below are the members who booked the conference")
	// for idx, usr := range members {
	// 	fmt.Printf("\nMember #%v \n", idx+1)
	// 	fmt.Println("Name : ", usr.Name)
	// 	fmt.Println("Email : ", usr.Email)
	// 	fmt.Println("Tickets Bought : ", usr.TicketsBought)
	// }
	return user, nil
}

func GetBookings() []models.User {
	fmt.Println("Getting all the bookings made so far")
	mu.Lock()
	defer mu.Unlock()
	result := make([]models.User, len(members))
	copy(result, members)
	return result
}

func UserDatabyEmail(email string) (models.User, error) {
	fmt.Println("Getting the User data mapped via Email")
	mu.Lock()
	defer mu.Unlock()
	user, ok := UserData[email]
	if !ok {
		return models.User{}, errors.New("User doesn't exist")
	}
	return user, nil
}

func UpdateTickets(usrEmail string, newTickets int) (models.User, error) {
	if newTickets <= 0 {
		return models.User{}, errors.New("Invalid new tickets , tickets must be > 0")
	}

	mu.Lock()
	defer mu.Unlock()

	oldData, ok := UserData[usrEmail]
	if !ok {
		return models.User{}, errors.New("Booking not found")
	}

	oldTickets := oldData.TicketsBought
	difference := newTickets - oldTickets

	if difference > RemainingTickets {
		return models.User{}, errors.New("Not enough tickets remaining")
	}

	oldData.TicketsBought = newTickets
	UserData[usrEmail] = oldData
	RemainingTickets -= difference

	for i := range members {
		if members[i].Email == usrEmail {
			members[i].TicketsBought = newTickets
			break
		}
	}

	return oldData, nil
}

func CancelBooking(usrEmail string) (models.User, error) {
	mu.Lock()
	defer mu.Unlock()

	usr, ok := UserData[usrEmail]
	if !ok {
		return models.User{}, errors.New("User doesn't Exist")
	}

	idx := -1
	for index, user := range members {
		if usrEmail == user.Email {
			idx = index
			break
		}
	}
	if idx == -1 {
		fmt.Println("Issue while finding the user , it exists in map but cant locate the index")
		return models.User{}, errors.New("User Exist , but cant locate the index")
	}
	tickets := usr.TicketsBought
	RemainingTickets += tickets
	delete(UserData, usrEmail)

	members[idx] = members[len(members)-1]
	members = members[:len(members)-1]

	return usr, nil
}

func GetTickets() int {
	mu.Lock()
	defer mu.Unlock()
	return RemainingTickets
}
