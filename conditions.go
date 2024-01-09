package main

import "fmt"

func main () {
	messageLen := 4
	maxMessageLen := 10


	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message Sent")
	}else {
		fmt.Println("Message not sent")
	}
}