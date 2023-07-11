package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	// This is a go routine which runs concurrently from the below code 
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func concurrency_test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	concurrency_test("Hello there Stacy!")
	concurrency_test("Hi there John!")
	concurrency_test("Hey there Jane!")
}