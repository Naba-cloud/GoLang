package main

import (
	"bufio"
	"fmt"
	"os"
)

// main is the entry point of the program.
// It welcomes the user, prompts for a pizza rating, and then thanks the user for their input.
func main() {
    // welcome is a string that holds the welcome message.
    welcome := "Welcome To The User Input"
    // Print the welcome message to the console.
    fmt.Println(welcome)
    
    // reader is a new buffered reader that reads from standard input (os.Stdin).
    reader := bufio.NewReader(os.Stdin)
    
    // Prompt the user to enter a rating for their pizza.
    fmt.Print("Enter The Rating For Your Pizza:")
    
    // Read the user's input until a newline character is encountered.
    // The second return value (error) is ignored using the blank identifier (_).
    input, _ := reader.ReadString('\n')
    
    // Thank the user for their rating and print the input received.
    fmt.Println("Thanks For Rating", input)
}
