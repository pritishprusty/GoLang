package main

import "fmt"

/* var arr [4]int //can be written as arr:= [4]int {1,2,3,4}

func main() {
	arr[0] = 1
	arr[1] = 3
	arr[2] = 4
	fmt.Println(arr) //we have initialized 3 values but we have 4 empty spcaes to fill so the output of last space is 0 for int
} */
//Creating an array whose size is determinded by the number of elements present in it
//Using Ellipsis
func main() {
	/* arr := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr) */ // Its dynamic accoding to input
	/* arr1 := [2][2]int{{1, 2}, {3, 5}} //two dimensional array
	fmt.Println(arr1) */

	/* myarr := [...]int{21, 3, 213, 43, 42, 5}
	//iterate using for loop
	for x := 0; x < len(myarr); x++ {
		fmt.Printf("%d\n", myarr[x])
	} */

	/* myarr := [...]int{21, 3, 213, 43, 42, 5}
	//iterate using range
	for index, value := range myarr {
		fmt.Println("element at", index, ":", value)
	} */

	/* arr2 := [2][2]int{{1, 2}, {3, 4}}
	for index1, value1 := range arr2 {
		for index2, value2 := range value1 { //the cycle continues and the 2nd value is used and explored
			fmt.Println("element at", index1, " ", index2, "is", value2)
		}
	} */
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr1 := arr[0:5] //from arr[0] to arr[4](excluding everything arr[5] onwards)
	arr1[0] = 10     //if we change the value here ,it'll change the original value as well because it takes pointer to change it
	fmt.Println(arr1)
	fmt.Println(arr)

}
