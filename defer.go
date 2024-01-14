// defer allows a function to be executed automatically just before its enclosing function returns.
// The deferred call's arguments are evaluated immediately, but the
// function call is not executed until the surrounding function returns.

// Deferred functions are typically used to close database connections, file handlers
// Defer is a great way to make sure that something happens at the end of a function,
// even if there are multiple return statements.

// defer keyword allow us to excute some function when the current function exist, or at the end of the current function

package main

import (
	"fmt"
	"sort"
)

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(users, name)
	user, ok := users[name]
	if !ok {
		delete(users, name)
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	delete(users, name)
	return logDeleted
}

// don't touch below this line

type user struct {
	name   string
	number int
	admin  bool
}

func test(users map[string]user, name string) {
	fmt.Printf("Attempting to delete %s...\n", name)
	defer fmt.Println("====================================")
	log := logAndDelete(users, name)
	fmt.Println("Log:", log)
}

func main() {
	users := map[string]user{
		"john": {
			name:   "john",
			number: 18965554631,
			admin:  true,
		},
		"elon": {
			name:   "elon",
			number: 19875556452,
			admin:  true,
		},
		"breanna": {
			name:   "breanna",
			number: 98575554231,
			admin:  false,
		},
		"kade": {
			name:   "kade",
			number: 10765557221,
			admin:  false,
		},
	}

	fmt.Println("Initial users:")
	usersSorted := []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")

	test(users, "john")
	test(users, "santa")
	test(users, "kade")

	fmt.Println("Final users:")
	usersSorted = []string{}
	for name := range users {
		usersSorted = append(usersSorted, name)
	}
	sort.Strings(usersSorted)
	for _, name := range usersSorted {
		fmt.Println(" -", name)
	}
	fmt.Println("====================================")
}
