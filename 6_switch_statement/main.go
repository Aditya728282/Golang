package main

import (
	"fmt"
	"time"
)

func main() {

	// switch statement in golang
	i := 6
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	case 4:
		fmt.Println("i is 4")
	case 5:
		fmt.Println("i is 5")
	default:
		fmt.Println("i is not in the range of 1 to 5")
	}

	// Multiple case in switch statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")

	default:
		fmt.Println("Workday")

	}
}
