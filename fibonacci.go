package main

import "fmt"

func main() {
	fmt.Println(fibo(6))

}
func fibo(number int) int {
	if number == 1 || number == 0 {
		return number
	}
	if number < 0 {
		fmt.Println("Invalid number")

		return -1
	}

	result := fibo(number-1) + fibo(number-2)
	return result
}
