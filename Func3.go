package main

import "fmt"

/*
func statement_name(){
	statement-1
	statement-2
	statement-3
}
*/

func main() {
	result, x := c()
	fmt.Println(result)
	fmt.Println(x)
	r, _ := b(1, 2, 3, 4, 5, 6)
	fmt.Println()
}
func a() (int, string) {
	// fmt.Print(12) // we can also use return statement to return something
	return 122, "Printing a"
}
func b(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}
func c() (i int, j string) {
	i = 10
	j = "returing from c"
	return

}
