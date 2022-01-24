package main

import "fmt"

//main func will be automatically called when you start the application
func main() { //place the curly braces right here after main function
	var data int
	data = 10
	// by giving : its not req to define the data type as we did above
	data1 := 1000
	data2 := "hello"
	// data:= 1000 is redeclaring data variables so will throw an error
	//data :=1000
	//if a var is not used then go will throw an error
	fmt.Print(data)
	fmt.Println(data1)
	fmt.Println(data2)
	// data types
	//1.primitive
	//int, float64, string, bool, complex,uint,byte( can be used to get unicode value)
	//2. non-primitive
	//struct, map, chan, interface,array, slice(like dynamic list array), rune, reflect, pointer
	// functional programming
	
}
