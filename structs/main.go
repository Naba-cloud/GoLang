package main

import "fmt"

// main is the entry point for the application.
// It demonstrates the creation and usage of a User struct.
func main() {
    fmt.Println("Welcome to Structs")
    
    // Create a new User instance with name, age, and email.
    naba := User{"Naba", 12, "naba@go.dev"}
    
    // Print the User instance using the default format.
    fmt.Println(naba)
    
    // Print specific fields of the User instance.
    fmt.Printf("Values of name %v and email is %v\n", naba.Name, naba.Email)
    
    // Print the User instance with field names.
    fmt.Printf("Values Are %+v\n", naba)
}
type User struct{
	Name string;
	Age int;
	Email string;


}