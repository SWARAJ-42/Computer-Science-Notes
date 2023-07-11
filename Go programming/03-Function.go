/*
	*** functions must be defined outside the main function
	*** function template
	func functionName(arg argtype, ...) returnType {
		code
		return returnTypeValue
	}

	*** func functionName(...) returnType ---> This is called a function signature
	*** if the all the args are of same type then write argtype only once
	*** Declaration syntax in go ---> f func(func(int,int) int, int) int
									  x int
									  p *int
									  a [3]int
	*** Go also uses the concept passbyValue and passbyReference
*/

package main

// calling multiple packages: no commas
import (
	"errors"
	"fmt"
)

// returns multiple valuesTypes
func fullName() (string, string) {
	return "Swaraj", "Biswal"
}
func subtract(num1 int, num2 int) int {
	return num1 - num2
}

func getCoords() (x, y int){ // x and y already initiaized
	// x and y are initialized with zero values
  
	return // automatically returns x and y but can be overrided by any other int variables by returning them explicitly
}

// Gaurded Clause --> returns a value or error earlier to maintain linearity
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero") // the function exits from this point
	}
	return dividend/divisor, nil  
}

func main() {
	fmt.Printf("%d\n",subtract(2, 3))
	firstName, _:= fullName() // underscore is used to ignore the second value
	fmt.Printf("%s\n", firstName) 
}