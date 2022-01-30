package main

import "fmt"

func SquareOfSum(n int) int {
	result := 0
	temp := 0
	for i := 1; i <= n; i++ {
		temp += i
	}
	result = temp * temp
	return result

}
func SumOfSquares(n int) int {
	result := 0
	temp := 0
	for i := 1; i <= n; i++ {
		temp += i * i
	}
	result = temp
	return result

}
func Difference(n int) int {
	var s1, s2, ret int
	s1 = SquareOfSum(n)
	s2 = SumOfSquares(n)
	ret = s1 - s2
	return ret

}
func main() {

	var n int
	fmt.Scanln(&n)
	fmt.Println(Difference(n))

}
