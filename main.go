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

		if userTickets < remainingTickets{
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
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		}
	}

}