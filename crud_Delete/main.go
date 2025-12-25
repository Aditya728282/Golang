// DELETE request CRUDE operation example in Go.

package main

import (
	"fmt"
	"net/http"
)

func performDeleteRequest() {

	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Delete request completed.")
	fmt.Println("Response Status:", resp.Status)
}

func main() {
	performDeleteRequest()
}
