package main

import (
	"fmt"
	"sort"
)

// main is the entry point of the program
func main() {
    fmt.Println("Welcome To Slices")

    // Slices Declaration Syntax
    var vegList = []string{"Apples", "Mango", "Peaches", "Tomato"}
    // In the above declaration, we don't specify the size; this is what is called a slice
    fmt.Printf("The type of vegList is: %T\n", vegList)

    // How to add values in slice: we cannot use the syntax we used in array vegList[0]=value
    // In Slices, we can use the append method; let's try it out
    vegList = append(vegList, "Dates", "Guava")

    // Let's see how we can print the values in a limit
    vegList = append(vegList[1:4])
    fmt.Println("fruit list with limit defined", vegList)

    fmt.Println("FruitsList", vegList)

    // How to add slices using make
    var highScores = make([]int, 4)
    highScores[0] = 9000
    highScores[1] = 200
    highScores[2] = 300
    highScores[3] = 400
    // But here we define the length of array so memory is allocated and we cannot allocate highScores[4]=34
    fmt.Println("HighScores", highScores)

    // When we use append, it will reallocate the memory
    highScores = append(highScores, 677, 899)
    fmt.Println("HighScores", highScores)

    // Sorting Slices
    sort.Ints(highScores)
    fmt.Println("HighScores", highScores)

    // Sorting Slices Of Strings
    sort.Strings(vegList)
    fmt.Println("Sorted Slices of strings", vegList)

    // Remove values from slices
}
