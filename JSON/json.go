// JSON: JavaScript Object Notation

package main

import (
	"encoding/json"
	"fmt"
)

// A simple struct to demonstrate JSON encoding/decoding
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {

	person := Person{Name: "Kallu", Age: 30, IsAdult: true}
	fmt.Println("Struct DATA: ", person)

	// Convert or Encoding struct to JSON (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("STRUCT to JSON: ", string(jsonData))

	// Convert or Decoding JSON to struct (Unmarshalling)
	var decodedPerson Person
	err = json.Unmarshal(jsonData, &decodedPerson)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("JSON TO STRUCT: ", decodedPerson)
}
