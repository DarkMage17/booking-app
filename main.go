package main

import (
	"booking-app/util"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go conference"
const conferenceTickets int = 50
var remainingTickets int = 50
var bookings = make([]UserData, 0)

type UserData struct{
	firstName string
	lastName string
	email string
	numberTickets int
}

var wg = sync.WaitGroup{}

func main(){

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	
	isValidName, isValidEmail, isValidTicketNumber := util.ValidateData(firstName, lastName, email, userTickets, remainingTickets)
	
	if isValidName && isValidEmail && isValidTicketNumber{
	
		bookTicket(userTickets, firstName, lastName, email)
	
		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)
	
		var firstNames = printFirstNames()
	
		fmt.Printf("These are all our bookings: %v\n", firstNames)
	
		if remainingTickets == 0 {
			fmt.Println("Tickets sold out!")
			//break
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
	wg.Wait()
}

func greetUsers(){
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func printFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames;
}



func getUserInput() (string, string, string, int){
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

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string){
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("#####################")
	wg.Done()
}