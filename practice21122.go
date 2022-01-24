package main

import "fmt"

func main() {
	C := 0
	for K, v := range "CAGTAA" {
		for J, l := range "TAGCCC" {
			if K == J && v-l == 0 {

				C++

			}

		}

	}
	fmt.Println("The hamming code is:", C)
}
