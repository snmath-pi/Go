package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend!")

	var bookings [50]string
	bookings[0] = "Saksham"
	var userName string
	var lastName string
	var email string
	var userTickets uint
	// ask the user info

	fmt.Println("Enter your firstName: ")
	fmt.Scan(&userName)
	fmt.Println("Enter your lastName: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter Number of tickets > 0: ")
	fmt.Scan(&userTickets)
	fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTickets, email)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)

}
