package main

import "fmt"

func main() {
	var Number uint // uint8 uint16 uint32 uint64 uintptr
	var isTrue bool
	var Number_float float64 // float32
	var String_var string
	var rune_value rune = 'a';
	default_declaration := "This is a default declaration"

	name := "Swaraj"
	var Greet string = fmt.Sprintf("Hello my name is %s\n", name)

	fmt.Printf("This is an %T with value %d.\n", Number, Number) // %T means printing the type of the variable
	fmt.Printf("This is an %T with value %v.\n", isTrue, isTrue) // %v means default
	fmt.Printf("This is an %T with value %f.\n", Number_float, Number_float)
	fmt.Printf("This is an %T with value %s.\n", String_var, String_var)
	fmt.Printf("This is an %T with value %c.\n", rune_value, rune_value)
	fmt.Printf("This is an %T with value %v.\n", default_declaration, default_declaration) // automatically detects the type for default datatypes
	fmt.Printf("%s", Greet)
}
