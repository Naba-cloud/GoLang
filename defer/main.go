package main

import "fmt"

// main is the entry point of the program.
// It demonstrates the use of the defer statement in Go.
func main() {
	// This deferred statement will be executed just before the function exits.
	// every time we use defer we can assume that
	// "this statament moves to the end of function no matter where it is written"
	defer fmt.Println("Hello")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	defer fmt.Println("Three")
	// Print a welcome message to the console.
	fmt.Println("Welcome To Defer In GOLang")

	// Print another message to the console.
	fmt.Println("World")
	DeferConcept()
}
func DeferConcept() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
