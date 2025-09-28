package main

import "fmt"

func main() {

	// for-> is use only for looping in  golang

	// while loop example
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// for loop exampl
	for i := 0; i <= 7; i++ {

		// break statement
		if i == 6 {
			break
		}

		// continue statement
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	// New Featured in Golang
	for i := range 10 {
		fmt.Println(i)
	}

	// Nested for loop
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i, j)
		}
	}

	// star pattern
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
