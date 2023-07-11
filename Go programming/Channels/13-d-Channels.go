package main 

import "fmt"

func channel_assign(iters int, ch_frange chan int) {
	for i:=0; i<iters;i++ {
		ch_frange <- i
	}
	close(ch_frange)
}

func main() {
	fmt.Println("Range in channels")
	ch_frange := make(chan int)
	go channel_assign(10, ch_frange)
	// here the value i iterated over range of values of ch_range
	for i := range ch_frange {
		fmt.Println(i)
	}
}