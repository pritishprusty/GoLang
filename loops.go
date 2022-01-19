package main

import "fmt"

/*
its for the second loop at line 18
for condition{
		statement
*/

func main() {
	// for i := 0; i < 10; i++ {
	// 	for i := 0; i < 3; i++ { // i is local here so we can take i in both loop as well. but its best to take another var
	// 		fmt.Println(i)
	// 	}

	// }
	// for true {
	// 	fmt.Println("infinite Execution")

	// }
	// i := 0
	// for i < 3 {
	// 	fmt.Println(i)
	// 	i++
	// }
	i := 0
	for {
		fmt.Println(i)

		if i == 0 {
			break
		}
	}
}
