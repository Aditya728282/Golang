package main

import (
	"fmt"
)

// Structure in golang

// personal data
type person struct {
	firstName string
	lastName  string
	age       int
}

// Education Details

type edu struct {
	highSchool string
	grad_      string
	course     string
}

// Contact Details
type contact struct {
	mob   int64
	email string
}

// Embede from other struct
// creating Employee structure

type employee struct {
	person
	edu
	contact
	position string
	salary   int64
}

func main() {

	fmt.Println("Structure in golang")

	// create an instance of person
	var person1 person
	person1.firstName = "john"
	person1.lastName = "doe"
	person1.age = 30

	// create an instance of education

	var edu1 edu
	edu1.highSchool = "ABC High School"
	edu1.grad_ = "XYZ University"
	edu1.course = "Computer Science"

	// create an instance of contact
	var contact1 contact
	contact1.mob = 78374659845
	contact1.email = "soyaa43@gmail.com"

	// create instANCE OF EMPLYEE
	var emp1 employee
	emp1.person = person1
	emp1.edu = edu1
	emp1.contact = contact1
	emp1.position = "Software Engineer"
	emp1.salary = 5436543

	fmt.Println("Employee Deatails")
	fmt.Println("contact Details:", emp1.contact)
	fmt.Println("Education Details:", emp1.edu)
	fmt.Println("Personal Details:", emp1.person)
	fmt.Println("Position:", emp1.position)
	fmt.Println("Salary:", emp1.salary)

	fmt.Println(emp1)

}
