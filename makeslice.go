// package main

// import (
// 	"errors"
// 	"fmt"
// )

// Most of the time we don't need to think about the underlying array of a slice.
//  We can create a new slice using the make function:
// func make([]T, len, cap) []T
// mySlice := make([]int, 5, 10)

// the capacity argument is usually omitted and defaults to the length
// mySlice := make([]int, 5)

//cap: capacity is the total space we have to grow the slice, lenght of the underline array
// len : lenght of the view into the array

// The length of a slice is simply the number of elements it contains.
// It is accessed using the built-in len() function:

// mySlice := []string{"I", "love", "go"}
// fmt.Println(len(mySlice)) // 3

// The capacity of a slice is the number of elements in the underlying array, counting
// from the first element in the slice.
// It is accessed using the built-in cap() function:

//  mySlice := []string{"I", "love", "go"}
//  fmt.Println(cap(mySlice)) // 3

package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	// ?
	costs := make([]float64, len(messages))

	for i := 0; i < len(messages); i++ {
			message := messages[i]
			cost := float64(len(message)) * 0.01
			costs[i] = cost
	}
	return costs
}

// don't edit below this line

func test(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Messages:")
	for i := 0; i < len(messages); i++ {
		fmt.Printf(" - %v\n", messages[i])
	}
	fmt.Println("Costs:")
	for i := 0; i < len(costs); i++ {
		fmt.Printf(" - %.2f\n", costs[i])
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test([]string{
		"Welcome to the movies!",
		"Enjoy your popcorn!",
		"Please don't talk during the movie!",
	})
	test([]string{
		"I don't want to be here anymore",
		"Can we go home?",
		"I'm hungry",
		"I'm bored",
	})
	test([]string{
		"Hello",
		"Hi",
		"Hey",
		"Hi there",
		"Hey there",
		"Hi there",
		"Hello there",
		"Hey there",
		"Hello there",
		"General Kenobi",
	})
}
