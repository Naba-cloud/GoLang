package main

import "fmt"

// main is the entry point of the program.
func main() {
    // Print a welcome message to the console.
    fmt.Println("Welcome to Pointers!")
    
    // Declare an integer variable and initialize it with a value.
    var myNumber = 23
    
    // Declare a pointer to an integer and assign it the address of myNumber.
    var ptr *int = &myNumber
    
    // Print the address stored in the pointer and the value of myNumber.
    fmt.Println("Address", ptr, myNumber)
    
    // Print the value pointed to by the pointer and the value of myNumber.
    fmt.Println("Number", *ptr, myNumber)
}
