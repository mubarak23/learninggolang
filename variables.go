// Variables in Go are passed by value (except for a few data types we haven't covered yet)
package main

import "fmt"

func main() {
	sendSoFar := 45
	const sendToAdd = 25
	incrementsends(sendSoFar, sendToAdd)
	fmt.Println("you've sent", sendSoFar, "messages")
}

func incrementsends(sendSoFar, sendToAdd int)  {
	sendSoFar = sendSoFar + sendToAdd
}

