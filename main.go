package main

import (
	"fmt"
	"sync"
	"time"
)

    var conferenceName="GO Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings =make([]UserData, 0)


	type UserData struct {
		firstName string
		lastName string
		email string
		numberOfTickets uint

	}
	var wg=sync.WaitGroup{}

 func main(){  
	greetUsers()
	
	//for {
	firstName, lastName, email, userTickets:=getUserInput();

	isValidName, isValidEmail, isValidTicketNumber:=ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)
   

	if  isValidName && isValidEmail && isValidTicketNumber {
		wg.Add(1)
		bookTickets(userTickets, firstName, lastName, email)
		go sendTicket(userTickets, firstName,lastName, email)
	
	firstNames:=getFirstNames()
    fmt.Printf("the first name of bookings are : %v\n",  firstNames)
	
	if remainingTickets==0 {
	//end the program
	fmt.Println("Conference is booked out, come back next year!")
	//break
    }
} else {
	if isValidName {
		fmt.Println("first or last name you entered is too short")
	}
	if !isValidEmail{
		fmt.Println("email address you entered doesn't contain @ sign")
	}
	if !isValidTicketNumber {
     fmt.Println( "number of tickets you enterd is invalid")
	}


}
	wg.Wait()
 //}
}


func greetUsers(){
	
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("conferenceTickets is %v, remainingTickets is %v, conferenceName is %v\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("we have a total of  %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets to attend\n")
}

func getFirstNames() []string{
	firstNames:=[]string{}
	for _,  booking:= range bookings{
		//var names=strings.Fields(booking)
		firstNames=append(firstNames, booking.firstName)

	}
   return firstNames
}



func getUserInput() (string, string, string, uint){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	
	fmt.Println("Enter your first name please")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name please")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email please")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets ypu want to book please")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string){
	remainingTickets=remainingTickets-userTickets
    
	//create map foe user
	var userData=UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
    
	bookings=append(bookings, userData)
	fmt.Printf("List of bookings is %v \n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. you will recieve a confirmation email at %v\n ", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remianing for %v\n", remainingTickets, conferenceName)
}

func  sendTicket(userTickets uint,firstName string, lastName string, email string){
	
     time.Sleep(10*time.Second)
	 ticket:=fmt.Sprintf("%v tickets for %v %v", userTickets,firstName, lastName)
	 fmt.Println("     ===============     ")
	 fmt.Printf("sending tickets:\n %v\n to email address %v\n", ticket, email)
	 fmt.Println(     "===============    ")

    wg.Done()
}