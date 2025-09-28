package main

import "fmt"

func main() {

	// if else conditions

	age := 300

	if age >= 18 {
		fmt.Println("person is an Adult")
	} else {
		fmt.Println("persont is not an  Adlut")
	}

	// if else if ladder
	if age < 10 {
		fmt.Println("person is a child")
	} else if age >= 10 && age < 20 {
		fmt.Println("person is a Teenager")
	} else if age >= 20 && age < 40 {
		fmt.Println("person is an Adult")
	} else if age > 40 && age < 70 {
		fmt.Println("Person is a Senior Citizen")
	} else {
		fmt.Println("Invalid age")
	}
}
