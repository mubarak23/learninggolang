// A receiver type on a method can be a pointer.

// Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver,
// pointer receivers are more common than value receivers.

package main

import "fmt"

// Pointer Recivers
type car struct {
	color string
}

type school struct {
	name string
}

func (s *school) setSchool (name string) {
	s.name = name
}

func (c *car) setColor(color string) {
	c.color = color
}

func main() {
	c := car{
		color: "white",
	}
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
	s := school {
		name : "Darius",
	}
	s.setSchool("Winners")
	fmt.Printf(s.name)
}
