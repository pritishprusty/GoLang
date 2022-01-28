package main

import "fmt"

func main() {
	//here our objective is to execute the f( ) at the end
	//our first option is to use f() at the end
	//our second choice is to use to defer

	fmt.Println(1)
	fmt.Println(2)
	defer f() //this will print the f( ) at the end of the main .
	fmt.Println(3)
	defer x() //multiple defer uses stack to operate to defer, as defer x goes at the end comes out at the end so
	//defer x() comes first then defer f()
	fmt.Println(4)
	//f()

}
func f() {
	fmt.Println("end statement")
}
func x() {
	fmt.Println("2nd defer")
}
