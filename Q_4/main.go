// Pattern Question

package main

import (
	"fmt"
)

func main() {

	//star prints a square pattern of stars
	n := 5
	fmt.Println("Square pattern of stars")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("* ")

		}
		fmt.Println()
	}
	fmt.Println()

	// number prints a square pattern of numbers
	fmt.Println("Square pattern of numbers")
	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//inceasing number pattern
	fmt.Println("Increasing number pattern")
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//decreasing number pattern
	fmt.Println("Decreasing number pattern")
	for i := 1; i <= n; i++ {
		for j := n; j >= 1; j-- {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//Alphabet pattern
	fmt.Println("Alphabet pattern")
	for i := 1; i < n; i++ {

		for j := 1; j <= n; j++ {
			fmt.Print(string(rune(j+64)), " ")
		}
		fmt.Println()
	}
	fmt.Println()

	//Reverse Alphabet pattern
	fmt.Println("Reverse Alphabet pattern")
	for i := 0; i < n; i++ {
		for j := n; j > 0; j-- {
			fmt.Print(string(rune(64+j)), " ")

		}
		fmt.Println()
	}

	// Square pattern of Number increasing order
	fmt.Println("Square pattern of Number increasing order")
	count := 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(count, " ")
			count++
		}
		fmt.Println()
	}

}
