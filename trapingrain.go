package main

import "fmt"

func main() {
	res := 0
	N :=4 
	var arr[]int ={3,0,2,4}

	temp := 0
	for i := 0; i <= N; i++ {
		for j := i + 1; j < N; j
		{
			if arr[i] > arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp

			}
		}
	}
	fmt.Println(arr[N-1])

}
