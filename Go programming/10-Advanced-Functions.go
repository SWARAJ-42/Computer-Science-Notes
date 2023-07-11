/*
	A higher-order function is a function that takes a function as an argument or returns a function as a return value. Go supports higher-order functions.

	Higher-order functions can be used for:-
		HTTP API handlers
		Pub/Sub handlers
		Onclick callbacks

	First-Class function - A first-class function is a function that can be treated like any other value. Go supports first-class functions. A function's type is dependent on the types of its parameters and return values.
	For example, these are different function types:
	func() int
	func(string) int

*/

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// this is an Higher-order function which takes another function as an arg to use it, note any function of the type mentioned can be used
func aggregator(a int, b int, c int, arithmetic func(int, int) int) int {
	return arithmetic(a, arithmetic(b, c))
}

// currying:- A function which takes a function as an argument and returns a new function as a value
func curr_add(x int, y int) int {
	return x + y
}
func curr_mul(x int, y int) int {
	return x * y
}

// This returns a new function called function currying
func selfMath(math_func func(int, int) int) func(int) int {
	return func(x int) int {
		return math_func(x, x)
	}
}

// Defer: The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns.

// Here is the following code snippet to understand defer
/*
	// CopyFile copies a file from srcName to dstName on the local filesystem.
	func CopyFile(dstName, srcName string) (written int64, err error) {

	// Open the source file
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// Close the source file when the CopyFile function returns
	defer src.Close()

	// Create the destination file
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	// Close the destination file when the CopyFile function returns
	defer dst.Close()

	return io.Copy(dst, src)
	}

*/

// Closures : A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.
// This concatter fucntion is a closure which stores doc as its reference variable for its own function body
func concatter() func(string) string {
	doc := "" // Here the doc variable will act as a global variable for all the concatter i.e no matter how many times the return typed function is run, it will never forget this variable.
	return func (str string) string { // This function never forgets doc variable and always updates it
		doc += str + " "
		return doc
	}
}

func main() {
	fmt.Println("Advanced functions")
	var a, b, c int = 3, 4, 6
	fmt.Println(aggregator(a, b, c, add)) // here "add" function is taken as arg because it satisfies the criteria of the defined argument

	// Curried function is declared
	Squared := selfMath(curr_mul) // Squared is now the type of function returned by selfMath(curr_mul)
	fmt.Println(Squared(5))

	// Closure functions in action 
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	//Output: Mr. and Mrs. Dursley of number four, Privet Drive

}
