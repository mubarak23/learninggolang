package main

import (
	"fmt"
	"strings"
)

// var hand int


func main () {
const	hand = 5
var myArray = [...]string{"first","second","third"}
// You can initialize the slice with values
mySlice := []string{"Fact","School","Bell"}
newSlice := append(mySlice, "Fouth", "Fiveth")

copySlice := make([]string, 3)

copy(newSlice, copySlice)
// Create slice of a specific length using the make function
var demoSlice = make([]string, 3)
copy(newSlice, demoSlice)

// map
//  initialize the map with out values 
ageMap := make(map[string]int)
ageMap["musa"] = 56

// You can also initialize the map with values directly using this syntax
demoAgeMap := map[string]int{"favoir": 45}

age := demoAgeMap["favoir"]

// delete an item from Map
delete(ageMap, "musa")

// loop in go with forloop
for i:=0; i < 10; i++ {
	fmt.Println(i)
}


	fmt.Println("Your hand number is ", hand)
	 msg := "Your hand number is " + fmt.Sprint(hand)
	 test := strings.HasPrefix("test", "te")
	 upper := strings.ToUpper("Welcome to Paris")
	fmt.Println(msg)
	fmt.Println(test)
	fmt.Println(upper)
	fmt.Println(myArray[1])
	fmt.Println(age)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	
}