package main

import "fmt"

/* Structure is just a blueprint
idea,it wont consume any of the ram until and unless the instace of it is created
Syntax-type struct_name struct{

}
*/

type Iphone12 struct {
	camera float32
	color  string
	ram    int
	rom    int
	price  float32
	size   float32
	//   call func () //we can define methods inside struc and outside its upto us but actuall implementation should be outside
	old Iphone13 //here the iphone13 struct is created outisde as a empty struct and when iphone12 is called this is called too

}
type Iphone13 struct {
	camera int
}

//this call function is exclusively available for the instance of iphone12 structure only
func (i Iphone12) call(number string) {
	fmt.Println("Calling..." + number)
}
func main() {
	//create instace of iphone
	//empty instance
	a := Iphone12{}    //an instance got created with default value for its attributes
	fmt.Println(a.old) //this is calling another struct from struct created outside the iphone12 i.e iphone13
	//this will print an empty {} as we didnt put any value in it
	a.camera = 12.4 //updating the camera attribute which priviously was 0
	a.color = "black"
	a.call("7504417346") //dot can be used to call var also functions

	fmt.Println(a)

	b := Iphone12{}
	fmt.Println(b)

	//non-empty instance
	c := Iphone12{
		camera: 10,
		ram:    12,
		rom:    12,
		size:   6,
		color:  "silver",
		price:  1,
		old: Iphone13{
			camera: 111,
		},
	} //put comma after 10 when putting a new line between 10 and }
	fmt.Println(c)

}
