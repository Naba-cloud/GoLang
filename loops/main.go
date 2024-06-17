package main

import (
    "fmt"
)

// main is the entry point of the program. It demonstrates various looping constructs in Go.
func main() {
    fmt.Println("Welcome To Loops Class")

    // days is a slice containing the days of the week.
    var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

    // Loop through the days slice using a traditional for loop with an index.
    for day := 0; day < len(days); day++ {
        fmt.Println(days[day])
    }

    // Loop through the days slice using a range loop, printing each day.
    for i := range days {
        fmt.Println("DAYS", days[i])
    }

    // Loop through the days slice using a range loop, printing the index and the day.
    for i, day := range days {
        fmt.Println(i, day)
    }

    // digit is an integer variable used in the following loop.
    digit := 0

    // Loop until digit is greater than 10.
    for digit <= 10 {
        if digit == 2 {
            digit++
            goto NC // Use goto to jump to the NC label.
        }
        if digit == 5 {
            digit++
            continue // Skip the rest of the loop and continue with the next iteration.
        }

        fmt.Println(digit)
        digit++
    }

NC:
    {
        // This block is executed when the goto statement jumps to the NC label.
        fmt.Println("Welcome to Naba Codes")
    }
}
