package main

import "fmt"

func main() {
	var a int = 0
	var b int = 1
	var n int
	fmt.Println("Enter the number of terms:")
	fmt.Scan(&n)
	fmt.Printf("Fibonacci Series up to %d terms:\n", n)

	for i := 0; i < n; i++ {
		next := a + b
		fmt.Printf("%d ", next)
		a = b
		b = next
	}

}
