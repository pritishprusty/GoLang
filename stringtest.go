package main

import (
	"fmt"
	"strconv"
)

func RunLengthEncode(input string) string {

	strcap := ""
	for i := 0; i < len(input); i++ {
		count := 1
		for i < len(input)-1 && input[i] == input[i+1] {
			count++
			i++
		}
		if count == 1 {
			strcap += string(input[i])
		} else {
			strcap += strconv.Itoa(count) + string(input[i])
		}
	}
	return strcap

}

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(RunLengthEncode(input))

}
