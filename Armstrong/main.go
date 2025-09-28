package main

import "fmt"

func main() {
	fmt.Println("Armstrong numbers")

	var n int
	fmt.Println("Enter a number:")
	fmt.Scan(&n)

	var temp int = n
	var sum int = 0

	for n != 0 {
		digit := n % 10
		sum = sum + digit*digit*digit
		n = n / 10
	}

	if temp == sum {
		fmt.Println("Armstrong number: ", temp)
	} else {
		fmt.Println("Not an Armstrong number")
	}
}
