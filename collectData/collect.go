package collectdata

import "fmt"

func CollectData() (string, string, int) {
	var name string
	var email string
	var tickets int

	fmt.Println("Enter your name")
	fmt.Scan(&name)
	fmt.Println("Enter your email")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets you want")
	fmt.Scan(&tickets)
	return name, email, tickets
}
