package main 

import "fmt" 

var pi float64 = 3.14159 

// interface defined 
type shape interface { 
	area() float64   // we can define arguments too if we want. for code beauty 
	perimeter() float64 
} 

// rectangle struct defined 
type rect struct { 
	length, breadth float64 
} 

// circle struct defined 
type circle struct { 
	radius float64 
} 

// area and perimeter corresponding to rectangle 
func (r rect) area() float64 { 
	return r.length * r.breadth 
} 
func (r rect) perimeter() float64 { 
	return 2*(r.length + r.breadth) 
} 

// area and perimeter corresponding to circle 
func (c circle) area() float64 { 
	return pi * c.radius * c.radius 
} 
func (c circle) perimeter() float64 { 
	return 2 * pi * (c.radius) 
} 

// interface as an argument in a function 
func findshape(s shape) { 
	fmt.Printf("The area of the shape is :%f\n", s.area()) 
	fmt.Printf("The perimeter of the shape is :%f\n", s.perimeter()) 
} 

func main() { 
	fmt.Println("Hello Interfaces") 
	// Interface implementation 
	r := rect{ 
		length: 12.44, 
		breadth: 13.33, 
	} 
	findshape(r) // during calling of the function interface is not used as argument but the struct 
} 