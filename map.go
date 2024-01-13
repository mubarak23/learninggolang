package main

import (
	"errors"
	"fmt"
)

// map provide key value mapping

// the zero value of a map is nil

// age := make(map([string][int]))
// age["John"] = 34

// demoAge := map[string][int]{
// 	"john": 35,
// 	"doe": 45
// }

// map mutations

// insert item inside a map
// m[key] = elem

// get an element from a map
// item := m[key]

// delete an item from a map
// delete(m, key)

// check if an item exist inside a map
// If key is in m, then ok is true. If not, ok is false.

// If key is not in the map, then elem is the zero value for the map's element type.

// elem, ok := m[key]




func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)

	if len(names) != len(phoneNumbers){
		return nil, errors.New("Invalid Size")
	}
	for i :=0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]

		userMap[name] = user {
			name,
			phoneNumber,
		}
	}
	return userMap, nil
}

// don't touch below this line

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func main() {
	test(
		[]string{"John", "Bob", "Jill"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	test(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)
}
