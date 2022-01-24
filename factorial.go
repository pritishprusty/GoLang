package main

import (
	"errors"
	"fmt"
) //do not put comma when importing nultiple files

func factorial(number int) (fact int, err error) { //this error is an interface/or can say a datatype {
	fact = 1
	if number < 0 {
		err = errors.New("number has to be posetive") //the errors here is the package
		return                                        //here we dont write the return var here because we go knows that the
		//the return var is fact and err because we specified it inside ()in line 8.

	} else if number == 0 {
		return

	}
	for number > 0 {
		fact *= number
		number--
	}
	return

}

func main() {
	var number int
	fmt.Println("enter Number")
	fmt.Scanln(&number)
	result, err := factorial(number)
	if err == nil {
		fmt.Println("the factorial is ", result)
	} else {
		fmt.Println(err)
	}

}
