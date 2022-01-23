package main

import (
	helper "booking-app/common"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const totalTickets uint = 50

var remainingTickets uint = totalTickets - 1

var bookings = []string{}

// var bookings = make([]map[string]string, 0)
// map : 1 single type for keys, 1 single type for values

var structBookings = []UserData{}

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// myArray := [3]string{}
// mySlice := []string{}
// myMap   := make(map[string]string)
// myStruct := StructType { key1: value1, key2: value2}

var bookingRequests = sync.WaitGroup{}

func main() {

	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidFirstName, isValidLastName, isValidEmail := helper.ValidateUserInput(firstName, lastName, email)

		if userTickets > remainingTickets {
			fmt.Println("Too many tickets")
			continue
		}

		if !isValidEmail || !isValidFirstName || !isValidLastName {
			fmt.Println("Invalid input")
			continue
		}

		bookTickets(firstName, lastName, email, userTickets)
		printFirstnames(bookings)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("No ticket left")
			break
		}
	}
	bookingRequests.Wait()
}

func greetUser() {
	fmt.Printf("conferenceName is %T and totalTickets %T\n", conferenceName, totalTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v are still available.\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here")
	stringTickets := strconv.FormatUint(uint64(totalTickets), 10)
	fmt.Printf("You know something useful about from/to string conversions ? strconv %v(%T) => %v(%T)\n", totalTickets, totalTickets, stringTickets, stringTickets)
}

func printFirstnames(bookings []string) {
	firstNames := getFirstnames(bookings)
	fmt.Printf("The first names: %v\n", firstNames)
}

func getFirstnames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask
	fmt.Print("Enter your firstName: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your lastName: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("How many tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName+" "+lastName)
	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The array type: %T\n", bookings)
	fmt.Printf("The array length: %v\n", len(bookings))

	fmt.Printf("User %v %v has %v tickets. Email sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets\n", remainingTickets)

	var userDataMap = make(map[string]string)
	userDataMap["firstName"] = firstName
	userDataMap["lastName"] = lastName
	userDataMap["email"] = email

	var userDataStruct = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	structBookings = append(structBookings, userDataStruct)

	fmt.Printf("firstname from map %v\n", userDataMap["firstName"])
	fmt.Printf("firstname from struct %v\n", userDataStruct.firstName)

	bookingRequests.Add(1)
	go sendTicket(userDataStruct)

}

func sendTicket(user UserData) {
	time.Sleep(10 * time.Second)
	fmt.Println("\n\n###################")
	fmt.Printf("sent tickets to user %v\n", user.email)
	fmt.Println("###################")
	bookingRequests.Done()
}
