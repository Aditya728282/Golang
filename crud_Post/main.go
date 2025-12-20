package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {

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

func performPostRequest() {

	// make data which you send to the database
	var todo = Todo{
		UserID:    1,
		Title:     "Aditya kumar",
		Completed: true,
	}

	// Always send data in JSON format to the server
	// convert todo struct to JSON format.
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// convert json data into string format
	jsonString := string(jsonData)
	fmt.Println("JSON Data: ", jsonString)

	// convert json string to Reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos/"

	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}

	// when respons is received then close the body
	defer res.Body.Close()

	// respone is response from the server
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response from server: ", string(data))

	fmt.Println("Response Status:", res.Status)
}

func main() {
	fmt.Println("Learning Crud POST Method")
	//performGetRequest()
	performPostRequest()
}
