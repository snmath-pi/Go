package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	// var bookings = [50]string{"Nana", "Nicole", "Peter"}
	// var bookings [50]string
	var bookings []string

	for {
		var userName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Println("Enter your firstName: ")
		fmt.Scan(&userName)
		fmt.Println("Enter your lastName: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter Number of tickets > 0: ")
		fmt.Scan(&userTickets)

		isValidName := len(userName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := userTickets <= remainingTickets
		isValid := isValidName && isValidEmail && isValidTicket
		if !isValid {
			if !isValidEmail {
				fmt.Println("Please Enter a valid Email (@)")
			}
			if !isValidName {
				fmt.Println("Please Enter a valid Name (name is too short)")
			}
			if !isValidTicket {
				fmt.Println("Please Enter a valid ticket amount")
			}

			continue
		} else {
			fmt.Printf("Thankyou %v %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, lastName, userTickets, email)
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userName+" "+lastName)
			fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("The first value: %T\n", bookings[0])
			fmt.Printf("Array Length: %v\n", len(bookings))

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)

			printfirstNames(bookings)
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our Conference is booked up. Come back next year!")
				break
			}
		}
	}

}

func greetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v\n", conferenceName)
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend!")
}

func printfirstNames(bookings []string) {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	fmt.Printf("The first Names of bookings are %v\n", firstNames)
}
