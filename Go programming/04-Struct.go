package main 

import "fmt"

// a struct
type SendMessageTo struct{
	PhoneNumber int
	message string
}

// passing a struct as an argument same as ny other datatype
func test(m SendMessageTo) {
	fmt.Printf("Sending message: %s\n", m.message)
	fmt.Printf("Sending message to: %d\n", m.PhoneNumber)
}

// nested Struct can be accessed by . operator 
type car struct {
	Make string
	Model string
	Height int
	Width int
	FrontWheel Wheel
	BackWheel Wheel
}
  
type Wheel struct {
	Radius int
	Material string
}

// main function
func main() {
	test(SendMessageTo{ // Ghost struct with no name and used only for once
		PhoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})

	// Initialization of a struct
	// there should be comma after every key value pair(Even the last)
	var Buggati car = car{
			Make: "NUll",
			Model: "Veyron",
			Height: 150,
			Width: 500,
			FrontWheel: Wheel{
				Radius: 20,
				Material: "Expensive Rubber",
			},
			BackWheel: Wheel{
				Radius: 20,
				Material: "Expensive Rubber",
			},
		}
	fmt.Printf("%s\n", Buggati.FrontWheel.Material)

	// anonymous structs initializing without any name to var
	myCar := struct { // we have declared the variable to an anonymous struct
		Make string
		Model string
	  } {
		Make: "tesla",
		Model: "model 3",
	}
	fmt.Printf("Make field: %s\n", myCar.Make)
	fmt.Printf("Model field: %s\n", myCar.Model)
	
	// anonymous struct inside a defined struct
	type car1 struct {
		Make string
		Model string
		Height int
		Width int
		// Wheel is a field containing an anonymous struct
		Wheel struct {
			Radius int
			Material string
		}
	}

	//Embedded structs .. more like of inheritance 
	type car2 struct {
		make string
		model string
	}
	  
	type truck struct {
	// "car" is embedded, so the definition of a
	// "truck" now also additionally contains all
	// of the fields of the car struct
		car2
		bedSize int
	}
	lanesTruck := truck{
		bedSize: 10,
		car2: car2{
			make: "toyota",
			model: "camry",
		},
	}
	
	fmt.Println(lanesTruck.bedSize)
	
	// embedded fields promoted to the top-level
	// instead of lanesTruck.car.make
	fmt.Println(lanesTruck.make)
	fmt.Println(lanesTruck.model)
}

/*
package main 

import "fmt"

// function defined for a struct
func (r rect) area() int {
	return r.length * r.breadth
}

type rect struct {
	length int
	breadth int
}

func main() {

	r2 := rect{
		length:12,
		breadth:13,
	}


	fmt.Println("%d", r2.area()) // run a function for a struct 
}
*/