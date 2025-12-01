// Create a file and write some content in file.

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("file.txt")
	// Check for errors during file creation
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Close the file when done ( free the os resources)
	defer file.Close()

	// Write some content to the file
	content := "Hey, I'm Aditya. This is a sample file created using Go."
	byte, writeErr := io.WriteString(file, content+"\n")

	fmt.Printf("Number of bytes written: %d\n", byte)

	if writeErr != nil {
		panic(writeErr)
	}

	fmt.Println("successfully Created file")
}
