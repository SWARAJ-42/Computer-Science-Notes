/*
	why generics ?
	- reduction of repetative code 
	
*/

package main 

import (
	"fmt"
)
// Here T is generic type; main point to note to define a generic function is to notice that square bracket where the generic is mentioned, this way the behaviour of the generic can be changed during the calling of the function 
func splitAnySlice[T any](s []T) ([]T, []T) {
    mid := len(s)/2
    return s[:mid], s[mid:]
}

func main() {
	var string_type []string = []string{
		"Swaraj",
		"Swadhin",
		"Joy",
		"Om",
		"Alice",
		"Bob",
	}
	var int_type []int = []int{
		1,2,3,4,5,6,
	}
	
	// Here it splits a slice of string type 
	left_group_string, right_group_string := splitAnySlice[string](string_type)
	fmt.Println("The left group:", left_group_string)
	fmt.Println("The right group:", right_group_string)

	// Here it splits a slice of int type
	left_group_int, right_group_int := splitAnySlice[int](int_type)
	fmt.Println("The left group:", left_group_int)
	fmt.Println("The right group:", right_group_int)
}