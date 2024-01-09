package main

import "fmt"

func concat (s1 string,  s2 string) string {
	return s1 + s2 
}

func main() {
	x := 4

	x = increment(x)
	fmt.Println(x)
	sendsSoFar := 430
	const sendsToAdd = 25
	incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")

	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go", " is fantastic")

// Ignore return values via _
	firstName, _ := getName()
	fmt.Println("Welcome to Textio,", firstName)
}

func increment (x int) int {
	 return x + 1
}

func test (s1 string, s2 string) {
		fmt.Println(concat(s1, s2))
}

func getName() (string, string) {
	return "Mubarak", "Aminu"
}

func incrementSends(sendSoFar, sendToAdd int) int {
	sendSoFar = sendSoFar + sendToAdd
	return sendSoFar
}

// Name return Values 
func getCords() (x, y int) {
	// 
	 // x and y are initialized with zero values

	 return // automatically returns x and y
	
}

// is the same as
func getCoords() (int, int){
  var x int
  var y int
  return x, y
}

// explicit return
// Even though a function has named return values, we can still explicitly return values if we want to.
func getCoordds() (x, y int){
  return x, y // this is explicit
}

// implicitly return  
func getCoordss() (x, y int){
  return // implicitly returns x and y
}



