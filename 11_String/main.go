package main

import (
	"fmt"
	"strings"
)

func main() {
	fruit := "aaple , Bnanana, orange ,mango"

	f := strings.Split(fruit, ",")
	fmt.Println(f)

	nam_1 := "aaditya"
	nam_2 := "Kumar"

	add := string.Join([]string{nam_1, nam_2}, " ")
	fmt.Println(add)

}
