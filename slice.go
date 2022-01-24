package main

import "fmt"

func main() {
	/* 	var s []int
	   	for index := range s {
	   		s = append(s, 20)

	   	}
	   	fmt.Println(s) */
	//syntax of make = make(data type,length ,capacity)
	//	s := make([]int, 5, 15) //5 here is the least space req and 15 is the max space it can have, we can also remove
	//15 here to make it dynamic
	var s []int //it is dynamic
	for i := 0; i < 6; i++ {
		s = append(s, 20) //this puts the value at the end
	}
	fmt.Println(s)
}
