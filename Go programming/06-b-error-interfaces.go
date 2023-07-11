package main

import (
	"errors"
	"fmt"
)

type userError struct {
	nameerr string
}

func (e userError) Error() string {
	return fmt.Sprintf("NameError: Username %s too large", e.nameerr)
}

func checkUsername(name string) error {
	if len(name) > 20 {
		// return userError{nameerr : name}
		// or using error package
		return errors.New(fmt.Sprintf("NameError: Username %s too large", name))
	}
	return nil
}

func main(){
	fmt.Println(checkUsername("WERTYIJHDSJHDSDVSjhdcsdcb"))
}