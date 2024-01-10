package main

import (
	"fmt"
)

// A convenient way to format strings in Go is by using the standard library's fmt.Sprintf() function.
// It's a string interpolation function, similar to JavaScript's built-in template literals.
// The %v substring uses the type's default formatting, which is often what you want.

func getSMSErrorString(cost float64, recipient string) string {
	// ?
	return	fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' can not be sent",
	 				cost, 
					recipient)
}

// don't edit below this line

func test(cost float64, recipient string) {
	s := getSMSErrorString(cost, recipient)
	fmt.Println(s)
	fmt.Println("====================================")
}

func main() {
	test(1.4, "+1 (435) 555 0923")
	test(2.1, "+2 (702) 555 3452")
	test(32.1, "+1 (801) 555 7456")
	test(14.4, "+1 (234) 555 6545")
}

