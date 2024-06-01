package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// main is the entry point of the program. It prompts the user to rate the Pizza App,
// reads the input, converts it to a float, and then adds 1 to the rating.
func main() {
    // fmt.Println("Welcome To Learning Conversions In GoLang")
    fmt.Println("Welcome To The Pizza App")
    fmt.Println("Please Rate Our Pizza App")

    // Create a new reader to read input from standard input (console)
    input := bufio.NewReader(os.Stdin)

    // Read the input string until a newline character is encountered
    reader, _ := input.ReadString('\n')
    fmt.Println("Thanks For Rating", reader)

    // Convert the input string to a float64 after trimming any whitespace
    numberRating, err := strconv.ParseFloat(strings.TrimSpace(reader), 64)
    if err != nil {
        // Print the error if the conversion fails
        fmt.Println(err)
    } else {
        // Print the incremented rating if the conversion is successful
        fmt.Println("Added 1 to Your Rating", numberRating+1)
    }
}
