package main

import "fmt"

// main is the entry point of the program.
func main() {
    fmt.Println("Welcome to Structs")

    // Creating a new User instance
    naba := User{"Naba", 12, "naba@go.dev", true}

    // Printing the User instance
    fmt.Println(naba)

    // Printing specific fields of the User instance
    fmt.Printf("Values of name %v and email is %v\n", naba.Name, naba.Email)

    // Printing the User instance with field names
    fmt.Printf("Values Are %+v\n", naba)
    
    // Calling methods on the User instance
    naba.GetStatus()
    naba.GetNewEmail()
}

type User struct {
	Name   string
	Age    int
	Email  string
	status bool
}

// GetStatus prints the active status of the User.
// It outputs a message indicating whether the user is active.
func (u User) GetStatus() {
    fmt.Println("Is User Active", u.status)
}
// GetNewEmail updates the User's email to a predefined address and prints the new email.
// This method sets the email to "naba@dev.io" and outputs the updated email to the console.
func (u User) GetNewEmail() {
    u.Email = "naba@dev.io"
    fmt.Println("Email is", u.Email)
}
