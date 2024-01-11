// slices are build on top of array

// A slice is a dynamically-sized, flexible view of the elements of an array.

// Slices always have an underlying array, though it isn't always specified explicitly.

// primes := [6]int{2, 3, 5, 7, 11, 13}
// mySlice := primes[1:4]
// mySlice = {3, 5, 7}

package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	// ?

	if plan == planPro {
		return allMessages[:], nil
	}
	if plan == planFree {
		return allMessages[0:2], nil
	}
	return nil, errors.New("unsuported plan")
}

// don't touch below this line

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func test(name string, doneAt int, plan string) {
	defer fmt.Println("=====================================")
	fmt.Printf("sending to %v...", name)
	fmt.Println()

	messages, err := getMessageWithRetriesForPlan(plan)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for i := 0; i < len(messages); i++ {
		msg := messages[i]
		fmt.Printf(`sending: "%v"`, msg)
		fmt.Println()
		if i == doneAt {
			fmt.Println("they responded!")
			break
		}
		if i == len(messages)-1 {
			fmt.Println("no response")
		}
	}
}

func main() {
	test("Ozgur", 3, planFree)
	test("Jeff", 3, planPro)
	test("Sally", 2, planPro)
	test("Sally", 3, "no plan")
}

