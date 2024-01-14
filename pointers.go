// pointer is a value that store a memory address
// pointer allow us to change the value of something from within a function

// The * syntax defines a pointer:
// var p *int
// the zero value for a pointer is nil

// The & operator generates a pointer to its operand.
// myString := "hello"
// myStringPtr := &myString
// & is use to reference a value
// the type of myString is a string
// the type of myStringPtr is a pointer to a string
// The * dereferences a pointer (to get underline value) to gain access to the value
// & to create new reference (pointers to a value) a value

// A receiver type on a method can be a pointer.

// Methods with pointer receivers can modify the value to which the receiver points.
//  Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

// Don't touch above this line

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

// Don't touch below this line

func test(recipient string, text string) {
	m := Message{Recipient: recipient, Text: text}
	sendMessage(m)
	fmt.Println("=====================================")
}

func main() {
	test("Lane", "Textio is getting better everyday!")
	test("Allan", "This pointer stuff is weird...")
	test("Tiffany", "What time will you be home for dinner?")
}
