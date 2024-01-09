package main

import "fmt"

func concat (s1 string,  s2 string) string {
	return s1 + s2 
}

func main() {
	x := 4

	x = increment(x)
	fmt.Println(x)

	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")
}

func increment (x int) int {
	 return x + 1
}

func test (s1 string, s2 string) {
		fmt.Println(concat(s1, s2))
}