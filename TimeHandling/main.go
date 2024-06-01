package main

import (
	"fmt"
	"time"
)

func statusUpdate() string { return "" }


// main is the entry point of the program. It demonstrates various functionalities of the time package.
func main() {
    // Print a welcome message
    fmt.Println("Welcome to learn about time packages")

    // Get the current time and print it
    presentTime := time.Now()
    fmt.Println(presentTime)

    // Format the current time in a specific layout and print it
    fmt.Println(presentTime.Format("02-02-2006 15:04 Monday"))

    // Create a specific date and time, format it, and print it
    createdDate := time.Date(2024, time.June, 3, 1, 0, 0, 0, time.UTC)
    fmt.Println(createdDate.Format("01-02-2006 Monday"))

    // Create a ticker that ticks every second and print its reference
    c := time.Tick(time.Second)
    fmt.Println("Time", c)
}
