package main

import (
	"fmt"
	"math"
	"time"
)


type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func(r rect) area() float64 {
	return r.width * r.height
}

func (c rect) perimeter() float64 {
	return 2*c.width + 2*c.height
}

type circle struct {
		raduis float64
}


func (c circle) area() float64 {
	return math.Pi * c.raduis * c.raduis
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.raduis
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	// ?
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("====================================\n")
}

func main() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "Bill Deer",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})
}

