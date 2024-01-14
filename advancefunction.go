// A programming language is said to have "first-class functions"
// when functions in that language are treated like any other variable

// a function can be passed as an argument to other functions, can be returned
// by another function and can be assigned as a value to a variable.

// A function that returns a function or accepts a function as input is called a Higher-Order Function.

// There are cases however when functions as values make a lot of sense. Some of these include:

// HTTP API handlers
// Pub/Sub handlers
// Onclick callbacks

// First class functions
// A first-class function is a function that can be treated like any other value. Go supports first-class functions.
// A function's type is dependent on the types of its parameters and return values.

// Definition: Higher-Order Functions

// A higher-order function is a function that takes a function as an argument or returns a function as a return value.
// Go supports higher-order functions. For example, this function takes a function as an argument:
// func aggregate(a, b, c int, arithmetic func(int, int) int) int

package main

import "fmt"

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

// don't touch below this line

func addSignature(message string) string {
	return message + " Kind regards."
}

func addGreeting(message string) string {
	return "Hello! " + message
}

func test(messages []string, formatter func(string) string) {
	defer fmt.Println("====================================")
	formattedMessages := getFormattedMessages(messages, formatter)
	if len(formattedMessages) != len(messages) {
		fmt.Println("The number of messages returned is incorrect.")
		return
	}
	for i, message := range messages {
		formatted := formattedMessages[i]
		fmt.Printf(" * %s -> %s\n", message, formatted)
	}
}

func add(x, y int) int {
  return x + y
}

func mul(x, y int) int {
  return x * y
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  return arithmetic(arithmetic(a, b), c)
}

// func main(){
//   fmt.Println(aggregate(2,3,4, add))
//   // prints 9
//   fmt.Println(aggregate(2,3,4, mul))
//   // prints 24
// }


func main() {
	fmt.Println(aggregate(2,3,4, add))
  // prints 9
  fmt.Println(aggregate(2,3,4, mul))
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addSignature)
	test([]string{
		"Thanks for getting back to me.",
		"Great to see you again.",
		"I would love to hang out this weekend.",
		"Got any hot stock tips?",
	}, addGreeting)
}
