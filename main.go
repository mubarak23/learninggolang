package main

import "fmt"

func concat (s1 string, s2 string) string {
	return s1 + s2
}
func main() {
	// name, age := "Demo", 45
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
	fmt.Println("Welcome to golang world")
}


func test (s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}