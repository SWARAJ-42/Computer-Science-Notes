/*
	Go programs express errors with error values. An error is any type that implements the simple built-in
*/

package main

import (
	"fmt"
)

func getUser() (string, error)  {
	user := ""
	return user, nil
}

func main(){
	fmt.Printf("Handling errors")
	user, err := getUser()
	if err != nil {
		fmt.Println(err)
		return // exit the program
	}
	// use user here
	fmt.Println(user)
}
