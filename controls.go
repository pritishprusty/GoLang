package main

import "fmt"

/*
1. if-else statement
if(condition){
	//statements
}

if(condistion){
	//statement
}else{
	staement
}
if(condition){
	//statements
}else{
	//statements
}
*/

func main() {
	fmt.Println("Enter a number:")
	var input int
	fmt.Scanln(&input)

	if input%2 == 0 {
		fmt.Println("hey you are even")
	} else {
		fmt.Println("hey you are odd")
	}

	if x := 10; x%2 == 0 {
		fmt.Println("even")

	} else {
		fmt.Println("odd")
	}

	/*
			2.switch(data){
			case var1:
				statement
		    case var2:
				statement
			default
			statement
			}

	*/

	data := 100
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
	case 100:
		data = 103
		fmt.Println(data)
		fallthrough
	case 11:
		data = 104
		fmt.Println(data)
		fallthrough
	case 15:
		data = 1002
		fmt.Println(data)
		fallthrough

	default:
		data = 10909
		fmt.Println(data)

	}

}
