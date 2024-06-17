package main

import (
	"fmt"
)

// main is the entry point of the program.
// It prints a welcome message, calls the greeter function,
// and then executes an immediately invoked function expression (IIFE)
// that prints another welcome message.
// and calls the Adder function to add two numbers.
// main is the entry point of the program. It demonstrates the usage of various functions.
func main() {
	fmt.Println("Welcome To Functions")
	greeter()

	// Anonymous function (Immediately Invoked Function Expression - IIFE)
	func() {
		fmt.Println("Welcome To IIFS")
	}()

	result := Adder(2, 3)
	fmt.Println("Result is", result)

	proSum := ProAdder(2, 3, 45, 6)
	fmt.Println("From ProAdder Functions", proSum)
	multiplicationResult, message := multiplier(3, 4)
	fmt.Println("messgae and multiplication result", multiplicationResult, message)
}

// greeter prints a greeting message to the console.
func greeter() {
	fmt.Println("Salam from greeter functions")
}

// Adder takes two integers and returns their sum.
//
// Parameters:
//
//	valOne (int): The first integer value.
//	valTwo (int): The second integer value.
//
// Returns:
//
//	int: The sum of valOne and valTwo.
func Adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// ProAdder takes a variable number of integer arguments and returns their sum.
//
// Parameters:
//
//	values: A variadic parameter representing the integers to be summed.
//
// Returns:
//
//	The sum of all the provided integer values.
func ProAdder(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
// multiplier takes a variadic number of integers and returns their product along with a descriptive string.
// 
// Parameters:
//   values (int...): A variadic number of integers to be multiplied.
//
// Returns:
//   int: The product of all the input integers.
//   string: A descriptive message indicating the result of the multiplication.
func multiplier(values ...int) (int, string) {
    mul := 1
    for _, value := range values {
        mul *= value
    }
    return mul, "Multiplication result from multiplier method"
}
