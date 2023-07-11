package main

import "fmt"

//If a pointer points to nothing (the zero value of the pointer type) then dereferencing it will cause a runtime error (a panic) that crashes the program. Generally speaking, whenever you're dealing with pointers you should check if it's nil before trying to dereference it.


// pointer reciever
type car_for_pt struct {
	color string
}
// the method is a pointer reciever, means it can change the original struct
// methods should be pointer recievers, functions must be value recievers
// Methods with pointer receivers don't require that a pointer is used to call the method. The pointer will automatically be derived from the value.
func (c *car_for_pt) setColor(color string) {
	c.color = color
}
func main(){
	x := 3
	fmt.Println("The value of x is : ", x)
	var p *int = &x // now p points to the address of x
	*p = 2 // the value of the address pointed by p is changed 
	fmt.Println("The value of x is : ", x)
	fmt.Println("The value of at the address pointed by p is : ", *p)


	c := car_for_pt{
		color: "white",
	}
	// notice here c is not a pointer 
	c.setColor("blue")
	fmt.Println(c.color)
	// prints "blue"
}