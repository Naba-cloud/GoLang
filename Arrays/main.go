package main

import "fmt"

// main is the entry point of the program.
func main() {
    // Print a welcome message to the console.
    fmt.Println("Welcome to Arrays")

    // Declare an array of strings with a fixed size of 4 to store fruit names.
    var fruitList [4]string

    // Assign values to the array elements.
    fruitList[0] = "Apple"
    fruitList[1] = "Mango"
    fruitList[3] = "Peach" // Note: fruitList[2] is left as the zero value (empty string).

    // Print the entire fruitList array to the console.
    fmt.Println(fruitList)

    // Print the length of the fruitList array to the console.
    fmt.Println(len(fruitList))

    // Declare and initialize an array of strings with a fixed size of 3 to store vegetable names.
    var vegList = [3]string{"Beans", "Potato", "Tomato"}

    // Print the entire vegList array to the console.
    fmt.Println(vegList)
}
