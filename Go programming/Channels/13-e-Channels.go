package main

import "fmt"

func main() {
	ch := make(chan int)
	go writeCh(ch)
	readCh(ch)
}

func readCh(ch <-chan int) {
	// ch can only be read from
	// in this function
	channel := <- ch
	fmt.Println(channel)
}

func writeCh(ch chan<- int) {
	// ch can only be written to
	// in this function
	ch <- 69
}
  