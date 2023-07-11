package main 

import "fmt"

func main(){
	fmt.Println("Lets go through maps")

	// Declaring a map
	ages := make(map[string]int) // here [string]int means keys are string type and values are int type
	ages["John"] = 12 // creates a new key-pair or we can say inserting a key-pair
	ages["Swaraj"] = 18
	ages["Swadhin"] = 19
	ages["Swadhin"] = 18 // overwrites the "Swadhin key-pair"
	delete(ages, "John") // deletes a particular key-pair
	// Check if a key-pair exist: below if a ok = True then the key exists and viceversa
	// Asking for a key that doesnot exist is a safe operation, you will be returned with a zero value
	// Asking for a value whose key doesnot exist is an unsafe operation, you will be returned with a panic
	value, ok := ages["John"]
	value_1, ok_1 := ages["Swadhin"]
	fmt.Println(value, ok)
	fmt.Println(value_1, ok_1)
	fmt.Println(ages)

	// Declaring and initializing maps at the same time
	new_ages := map[string]int{"Swaraj": 19, "Swadhin":19}
	fmt.Println(new_ages)

	// Struct vs map as maps
	// it is better to use Structs over maps as values because it becomes tedious and more prone to panic when we use maps as values

}