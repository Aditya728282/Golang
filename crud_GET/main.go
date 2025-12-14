package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning Crud GET Method")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Response Status:", resp.Status)
	}

	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Todo: ", todo)

}
