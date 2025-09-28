// Palindrome number program in Go

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scan(&n)

	temp := n
	rev := 0
	for n != 0 {
		digit := n % 10
		rev = rev*10 + digit
		n = n / 10

	}
	fmt.Println(rev)

	if rev == temp {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}
}
