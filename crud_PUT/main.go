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

// UpdateRequest
func performUpdateRequest() {

	// Update or Replace data which you want to send to the database

	var todo = Todo{
		UserID:    3010,
		Title:     "Aditya Kumar Singh In Capgemini",
		Completed: false,
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
	const myURL = "https://jsonplaceholder.typicode.com/todos/1"

	// There is no http.Put method is available like Get and Post
	// So we have to create our own request using http.NewRequest method
	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		return
	}

	//set the content-type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// send the request using http client and store the response
	cleint := http.Client{}
	res, err := cleint.Do(req)
	if err != nil {
		fmt.Println("Error making PUT request:", err)
		return
	}

	// when respons is received then close the body
	defer res.Body.Close()

	// respone is reader from the server
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response from server: ", string(data))

	fmt.Println("Response Status:", res.Status)

}

func main() {
	fmt.Println("Learning Crud Update Method")
	performUpdateRequest()
}
