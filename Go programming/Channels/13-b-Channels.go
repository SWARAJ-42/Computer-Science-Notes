package main

import (
	"fmt"
	"time"
)

func filterOldEmails(emails []email) {
	// defining a channel
	isOldChan := make(chan bool)
	// defining a channel of a buffer length
	// Sending on a buffered channel only blocks when the buffer is full.
	// Receiving blocks only when the buffer is empty.
	ch := make(chan int, 100)
	close(ch) // closing a channel
	_, ok := <- ch // To check whether the channel is closed, here ok is the boolean to check and _ is the value of ch
	if !ok {
		fmt.Println("Channel closed !!!")
	}
	go func(isOldChan chan<- bool, emails []email) {
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {    
				isOldChan <- true   // as soon as this or the latter is assigned line 20-25 is run and a new go routine starts to fill the channel(isOldChan) again otherwise the program will halt if is "Oldchain" is empty                    
				continue                                                	   
			}
			isOldChan <- false
		}
	}(isOldChan, emails)

	isOld := <-isOldChan
	fmt.Println("email 1 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old:", isOld)
}

type email struct {
	body string
	date time.Time
}

func channel_test(emails []email) {
	filterOldEmails(emails)
	fmt.Println("==========================================")
}

func main() {
	channel_test([]email{
		{
			body: "Are you going to make it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I need a break",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What were you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	})
}