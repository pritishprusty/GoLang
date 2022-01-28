package main

import "fmt"

func main() {
	devide(4, 0)
}

/* func devide(a, b int) int {
	defer recoverB() {
		fmt.Println(recover()) //recovery function creates helps to handle error but it needs a function to operate
		//hence we use a nameless funciton to do this
	}()
	return a / b
} */

func devide(a, b int) int {
	defer recoverB()
	if b == 0 {
		panic("b is the issue") //this is custom panic ,
	}
	return a / b
}

func recoverB() {
	if r := recover(); r != nil {
		fmt.Println("recovered")
	}
}

//if recover is not there ,panic will happen then termination of program will happen so we need this recover to handle the panic
