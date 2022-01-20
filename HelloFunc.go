package main

import (
	"errors"
	"fmt"
)

func greetings(name string) (string, error) { //the last string is the return type
	if name == "" {
		return name, errors.New("Name can not be Null")
	}
	return "Hello," + name, nil
}
func main() {
	greetings, err := greetings("Pritish")
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(greetings)
	}
}
