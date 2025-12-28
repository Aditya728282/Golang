package main

import (
	"fmt"
	"time"
)

func MyName() {
	fmt.Println("My First Name is Aditya Kumar Singh")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("My Phone Number is 1234567890")
}

func Education() {
	fmt.Println("I have studied in XYZ College")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("I'm a Software Engineer")
}

func main() {
	fmt.Println("Hello, World!")
	go MyName()
	go Education()
	// time.Sleep(1000 * time.Millisecond)
	time.Sleep(3000 * time.Millisecond)
}
