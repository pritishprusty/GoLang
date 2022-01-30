package main

import (
	"fmt"
	"strconv"
)

func RunLengthEncode(input string) string {

	n := len(input)
	StringStore := ""
	for i := 0; i < n; i++ {
		count := 1
		for i < n-1 && input[i] == input[i+1] {
			count++
			i++
		}
		if count == 1 {
			StringStore += string(input[i])
		} else {
			StringStore += strconv.Itoa(count) + string(input[i])
		}
	}
	return StringStore

	// var ret int
	// ret = 0
	// var ret2 string
	// var ret4 string
	// for i := 0; i < len(input); i++ {
	// 	if input[i] == input[i+1] {
	// 		ret = ret + 1
	// 		ret2 = strconv.Itoa(ret) + string(input[i])

	// 	} else {
	// 		ret2 = "%s ,input[i]"
	// 	}
	// }
	// ret4 += ret2
	// return ret4
}

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(RunLengthEncode(input))

}
