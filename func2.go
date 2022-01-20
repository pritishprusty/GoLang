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

func add(num ...int) int { // the three dots here represents multiple values that can go into num and its dynamic (like list or array but it does not has index)
	sum := 0
	for _, value := range num { // use _, to make the program right,this used in place of index,so if we dont write index we use _,
		sum = sum + value // if we dont use _, it wont count the last  value
	}
	return sum

}

func main() {
	greetings, err := greetings("Pritish")
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(greetings)
	}

	fmt.Println("the sum of 12345 is ", add(1, 2, 3, 4, 5))
	fmt.Println("the sum of 1,2,3,4,5,6", add(1, 2, 3, 4, 5, 6))

}
