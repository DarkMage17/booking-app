package main

import (
	"fmt"
	"strings"
)

func main(){
	var conferenceName string = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings[] string

	for remainingTickets > 0 && len(bookings) < 50{
		var firstName string
		var lastName string
		var email string
		var userTickets int
	
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber{
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)
		
			fmt.Printf("The whole array: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
		
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v", remainingTickets, conferenceName)
		
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
	
			fmt.Printf("These are all our bookings: %v\n", firstNames)
	
			if remainingTickets == 0 {
				fmt.Println("Tickets sold out!")
				break
			}
		}else {
			if !isValidName{
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("Email address doesn't contain @")
			}
			if !isValidTicketNumber{
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}

}