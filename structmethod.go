package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func (p Person) Speak() {
	fmt.Println("Hello From ", p.Name)
}

func main () {
	demo := Person{ Age: 34, Name: "Flavoi"}


	demo.Speak()
}