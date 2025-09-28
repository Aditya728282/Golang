package main

import (
	"fmt"
)

func main() {
	// Create a map
	student := make(map[string]int)

	// Insert key-value pairs
	student["Alice"] = 21
	student["Bob"] = 22
	student["Charlie"] = 23

	// Add a new key-value pair
	student["David"] = 24
	fmt.Println("After adding David:", student)

	// Update an existing key-value pair
	student["Alice"] = 25
	fmt.Println("After updating Alice's age:", student)

	// Delete a key-value pair
	delete(student, "Bob")
	fmt.Println("After deleting Bob:", student)

	for k, v := range student {
		fmt.Printf("%s : %d\n", k, v)
	}

}
