package main

import "fmt"

type messageSend struct {
	phoneNumber int
	message string
}

type car struct {
	Make string
	Model string
	Height string
	Width string 
	FrontWheel Wheel
	BackWheel Wheel
}

type Wheel struct {
	Radius int
	Material string
}

type messageToSend struct {
	message string
	sender  user
	recipient user
}


type user struct {
	name string
	number int
}

func canSendMessage (mToSend messageToSend) bool {
	if mToSend.recipient.name == ""{
		return false
	}

	if mToSend.recipient.number == 0 {
		return false
	}

	if mToSend.sender.name == "" {
		return false
	}

	if mToSend.sender.number == 0 {
		return false
	}
	return true
}


func demoTest (toSend messageToSend) {
	fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
				toSend.message,
				toSend.sender.name,
				toSend.sender.number,
				toSend.recipient.name,
				toSend.recipient.number,
		)
		fmt.Println()

		if canSendMessage(toSend) {
			fmt.Println("...sent!")
		} else {
			fmt.Println("...can't send message")
		}

		fmt.Println("====================================")

}

func test (m messageSend) {
	fmt.Printf("Sending message '%s' to: %v\n", m.phoneNumber, m.message)

	fmt.Println("====================================")

}


func main() {
	test(messageSend{
		phoneNumber: 234567945,
		message: "Love to have you aboard!",
	})

	test(messageSend{
		phoneNumber: 9804583734,
		message: "Thank you for signup",
	})

	test(messageSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})
	myCar := car{}
	myCar.FrontWheel.Radius = 5 // We use dot operator to access field (.) 
	fmt.Println(myCar)

	demoTest(messageToSend{
		message: "you have an appointment tommorow",
		sender: user{
			name:   "Brenda Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	demoTest(messageToSend{
		message: "you have an event tommorow",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "Suzie Sall",
			number: 19035558973,
		},
	})
	demoTest(messageToSend{
		message: "you have an party tommorow",
		sender: user{
			name: "Njorn Halafax",
		},
		recipient: user{
			name:   "Becky Sue",
			number: 19035558973,
		},
	})
	demoTest(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Eli Halafax",
			number: 16545550987,
		},
		recipient: user{
			number: 19035558973,
		},
	})
	demoTest(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Jason Bjorn",
			number: 16545550987,
		},
		recipient: user{
			name: "Jim Bond",
		},
	})
	
}