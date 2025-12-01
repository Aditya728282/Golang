// URL = Uniform Resource Locator

package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL Parsing Example")
	myurl := "https://www.example.com:8080/path/to/resources?key1=value1&key2=value2query=123"
	fmt.Printf("Type of myurl: %T\n", myurl)

	// It is used to convert string to URL object
	parsedURL, err := url.Parse(myurl)
	fmt.Printf("After parsing type of Url is : %T\n", parsedURL)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	/* Accessing URL COMPONENTS:
	1. Scheme: https, http, ftp
	2. Host: www.example.com
	3. Path: /path/to/resources
	4. RawQuery: key1=value1&key2=value2query=123
	5. Port: 8080
	7. User: username:password
	*/

	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host: ", parsedURL.Host)
	fmt.Println("Path: ", parsedURL.Path)
	fmt.Print("RawQuery:", parsedURL.RawQuery)

	// Modify URL COMPONENTS

	parsedURL.Scheme = "http"
	parsedURL.Host = "www.changedexample.com"
	parsedURL.Path = "/new/path"
	parsedURL.RawQuery = "newkey=newvalue"

	newUrl := parsedURL.String()
	fmt.Println("\nModified URL:", newUrl)
}
