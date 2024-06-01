package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main is the entry point of the program.
func main() {
	// Create a new random source seeded with the current time in nanoseconds.
	source := rand.NewSource(time.Now().UnixNano())

	// Create a new random number generator using the source.
	rng := rand.New(source)

	// Generate a random integer between 1 and 6 (inclusive).
	Dice := rng.Intn(6) + 1

	// Print the generated random integer.

	switch Dice {
	case 1:
		fmt.Println("Dice Value 1")
	case 2:
		fmt.Println("Dice Value 2")
		fallthrough
	case 3:
		fmt.Println("Dice Value 3")
	case 4:
		fmt.Println("Dice Value 4")
	case 5:
		fmt.Println("Dice Value 5")
	case 6:
		fmt.Println("Dice Value 6")
	default:
		fmt.Println("what was that?")
	}
}
