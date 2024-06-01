package main

import (
    "fmt"
)

// main is the entry point of the program. It demonstrates how to remove a value from a slice.
func main() {
    fmt.Println("How to remove values from Slices")
    
    // Initialize a slice of programming languages
    var languages = []string{"Javascript", "GO", "Python", "Java", "Rust"}
    
    // Index of the element to be removed (Python in this case)
    var index int = 2
    
    // Remove the element at the specified index
    languages = append(languages[:index], languages[index+1:]...)
    
    // Print the updated slice
    fmt.Println("languages", languages)
}
