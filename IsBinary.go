// Check if a number is binary

package main

import "fmt"

// Returns true if the number contains only 0s and 1s
func isBinary(num int) bool {
	if num < 0 {
		return false
	}
	if num == 0 {
		return true
	}
	for num != 0 {
		digit := num % 10
		if digit > 1 {
			return false
		}
		num = num / 10
	}
	return true
}

func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scan(&n)
	res := isBinary(n)
	fmt.Println(res)
}
