package main

import "fmt"

func main() {
	// star prints a right angled triangle pattern of stars
	n := 5
	fmt.Println("Right triangle pattern of stars")
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
	fmt.Println()

	// number prints a right angled triangle pattern of numbers
	fmt.Println("Right triangle pattern of numbers")
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}

	fmt.Println()

	// Alphabet pattern triangle
	fmt.Println("Alphabet pattern triangle")
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(string(rune(j+65)), " ")
		}
		fmt.Println()
	}

}
