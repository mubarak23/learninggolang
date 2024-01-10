package main

import "fmt"

// Because errors are just interfaces, you can build your own custom types that implement the error interface.
// Here's an example of a userError struct that implements the error interface:

type userError struct {
	name string
}
// we can build our type that implement the Errror Interface


// our type that implement the Errror Interface 
func (e userError ) Error() string {
	 return fmt.Sprintf("%v has a problem with their account", e.name)
}

// func sendSMS(msg, userName string) error {
// 	if !canSendToUser(userName) {
// 			return userError{name: userName}
// 	}
// 	...
// }


// Update the code so that the divideError type implements the error interface. 
// Its Error() method should just return a string formatted in the following way:
// can not divide DIVIDEND by zero


type divideError struct {
	dividend float64
}

// using the divideError Struck to implement the Error interface

func (de divideError) Error() string {
		return fmt.Sprintf(
			"can not divide %v by zero", 
			de.dividend,
		)
}

// don't edit below this line

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		// We convert the `divideError` struct to an `error` type by returning it
		// as an error. As an error type, when it's printed its default value
		// will be the result of the Error() method
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func test(dividend, divisor float64) {
	defer fmt.Println("====================================")
	fmt.Printf("Dividing %.2f by %.2f ...\n", dividend, divisor)
	quotient, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Quotient: %.2f\n", quotient)
}

func main() {
	test(10, 0)
	test(10, 2)
	test(15, 30)
	test(6, 3)
}


