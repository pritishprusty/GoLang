package main

import (
	"fmt"
	"reflect"
)

//main func will be automatically called when you start the application
func main() { //place the curly braces right here after main function
	// 	var data int
	// 	data = 10
	// 	// by giving : its not req to define the data type as we did above
	// 	data1 := 1000
	// 	data2 := "hello"
	// 	// data:= 1000 is redeclaring data variables so will throw an error
	// 	//data :=1000
	// 	//if a var is not used then go will throw an error
	// 	//1.primitive
	// 	//int, float64, string, bool, complex,uint,byte( can be used to get unicode value)
	// 	//2. non-primitive
	// 	//struct, map, chan, interface,array, slice(like dynamic list array), rune, reflect, pointer
	// 	// functional programming
	/* const pi flot32 = 3.14 //explicit typing
	   //const pi=3.13 //implicit typing

	   const l int // we can not write const in two lines
	   l =30 //like here
	   //const is constant */
	/* 	const pi float32 = 3.14 //explicit typing
	   	//const pi=3.13 //implicit typing

	   	const l int = 100
	   	fmt.Println(reflect.TypeOf(pi)) //reflext.typeof will print the data type of the variaable
	   	fmt.Println(l) */
	var u rune
	u = 'a'
	fmt.Println(u)                 //rune takes the ascii value and prints it
	fmt.Println(reflect.TypeOf(u)) //rune a 4 bytes and is of int32

}
