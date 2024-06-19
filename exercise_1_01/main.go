package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println(stars)

	var total float64 = 2 * 13
	fmt.Println("Sub :", total)

	total = total + (4 * 2.25)
	fmt.Println("Sub :", total)

	total = total - 5
	fmt.Println("Sub :", total)

	tip := total * 0.1
	fmt.Println("Tip :", tip)

	split := total / 2
	fmt.Println("Split:", split)

	visitCount := 24
	visitCount = visitCount + 1

	remainder := visitCount % 5

	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward.")
	}

	var count int
	fmt.Printf("Count  : %#v \n", count)

	var discount float64
	fmt.Printf("Discount : %#v \n", discount)

	var debug bool
	fmt.Printf("Debug  : %#v \n", debug)

	var message string
	fmt.Printf("Message : %#v \n", message)

	var emails []string
	fmt.Printf("Emails : %#v \n", emails)

	var startTime time.Time
	fmt.Printf("Start  : %#v \n", startTime)
}
