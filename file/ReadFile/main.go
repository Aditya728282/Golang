package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	/*
		file, err := os.Open("file.txt")

		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		// Create a Buffer to hold file content(store data for temporary use)
		buffer := make([]byte, 1024)

		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			// Print the content read from the file
			fmt.Print(string(buffer[0:n]))
		}
	*/

	content, err := ioutil.ReadFile("file.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(string(content))
	fmt.Println("File read successfully")
}
