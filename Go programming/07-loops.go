package main

import "fmt"

func main() {
	fmt.Println("Lets loop with loops")
	// for loop in go is almost similar to c-for-loops but without brackets in on the conditionals
	// remove the middle condition and the loop runs forever
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("")
	// There are no while loops in go instead we can use for loops for this 
	i := 0
	for i < 5 {
		fmt.Printf("%d ", i+1)
		i++
	}
	fmt.Println("")

	// break and continue are same as python
}