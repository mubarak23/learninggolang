package main

import "fmt"


func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// simulate while loop
	i := 0

for i < 10 {
	fmt.Println(i)
  i++
}

t := 0

for { 
	fmt.Println(t)

	if t < 10 {
		break

	}
}

numbers := []int{1, 2, 3}

for i, num := range numbers {
	fmt.Printf("%d: %d\n", i, num)
}

	
}