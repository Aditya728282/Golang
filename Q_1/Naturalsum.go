// Sum of sum-series of first N Natural numbers

package main

import "fmt"

func NaturalSum(n int) int {

	var sum int = 0
	var i int
	for i = 1; i <= n; i++ {
		sum = sum + (i*(i+1))/2
	}
	return sum
}

func main() {

	fmt.Println("Enter a size of N")

	var res int = NaturalSum(5)
	fmt.Println(res)
}
