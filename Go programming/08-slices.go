package main 

import "fmt"

func main() {
	// Unlike JS arrays go arrays are fixed and we have to predefine the size of array 

	// var myInts [10]int // array of 10 zeros
	primes := [6]int{2,3,5,7,11,13}

	// These are the ways to print the array 
	fmt.Printf("%d\n", primes)
	fmt.Printf("%+v\n", primes)
    fmt.Printf("%#v\n", primes)
    fmt.Printf("%+q\n", primes)

	// slices are to extract sub value arrays from a parent array 
	// unlike arrays slices need not be of fixed size 
	// slices and array are passbyreference not passbyvalue
	slice := primes[1:3] // excluding index at 3 , we can modify this new array becausenow its a slice
	fmt.Println(slice)

	// explicitly creating a slice for dynamic memory allocation 
	// make([]type, length, capacity{defaults to length})
	mySlice := make([]int, 5, 10) // 10 can be omitted for the capacity to be 5 and expandable(reallocating the array to a new location and doubling it)
	fmt.Printf("%d\n",mySlice)
	mySliceLiterals := []string{"I", "Love", "Go"}
	fmt.Printf("%s\n",mySliceLiterals)
	
	// variadic functions : function with slice as its args, can be of any length 
	// Example of some built in variadic function are Sprintf and Printf
	// func Println(a ...interface{}) (n int, err error)
	fmt.Println("The sum is : ",addInt(1,2))
	fmt.Println("The sum is : ",addInt(1,2,3))
	fmt.Println("The sum is : ",addInt(1,2,3,4))
	fmt.Println("The sum is : ",addInt(1,2,3,4,5))
	// Spread operator : used for pass spreading the slices into separate values
	nums := []int{1,2,3,4,5,6}
	fmt.Println("The sum is : ", addInt(nums...)) // here ... used after the var is the spread operator

	// Append function
	// func append(slice []Type, elems ...Type) []Type
	// syntax: slice = append(slice, onething, secondThing, .......) 
	// append syntax manages everything and you dont have to worry about memory reallocation or capacity if performance is not the concern
	nums = append(nums, 7) // here 7 is appended as the last element 

	// 2D slices 
	matrix := make([][]int, 0)
	// making a 10 x 10 matrix
	for i := 0; i < 10; i++ {
		row := make([]int, 10)
		for j := 0; j < 10; j++ {
			row[j] = j * i 
		}
		matrix = append(matrix, row)
	}

}

// variadic function def
func addInt(num ...int) int {
	sum:=0
	for i:=0; i < len(num); i++{
		sum += num[i]
	}
	return sum
}