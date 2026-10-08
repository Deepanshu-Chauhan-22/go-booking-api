package models

type User struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	TicketsBought int    `json:"tickets"`
}

//var Wg sync.WaitGroup
