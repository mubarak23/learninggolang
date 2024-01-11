package main

import "fmt"

// Many functions, especially those in the standard library,
// can take an arbitrary number of final arguments.
// This is accomplished by using the "..." syntax in the function signature.

// A variadic function receives the variadic arguments as a slice.

func concat(strs ... string) string {
	final := ""
	for _, str := range strs {
		final += str
	}
	return final
}

func printStrings(strings ...string) {
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}
}

// func main() {
//     names := []string{"bob", "sue", "alice"}
//     printStrings(names...)
// }


func sum(nums ...float64) float64 {
	total := 0.0
	for i := 0; i < len(nums); i++ {
			total += nums[i]
	}
	return total
}

// // don't edit below this line

func test(nums ...float64) {
	total := sum(nums...)
	fmt.Printf("Summing %v costs...\n", len(nums))
	fmt.Printf("Bill for the month: %.2f\n", total)
	fmt.Println("===== END REPORT =====")
}

func main() {
	final := concat("Hello", "World", "Friends")
	fmt.Println(final)
	names := []string{"bob", "sue", "alice"}
    printStrings(names...)
	test(1.0, 2.0, 3.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0)
	test(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0)
}
