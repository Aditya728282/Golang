//  Lowercase to Uppercase
// aditykumar convert into ADITYAKUMAR

package main

import "fmt"

func main() {
	name := "aditykumar"

	for i := 0; i < len(name); i++ {
		if name[i] >= 'a' && name[i] <= 'z' {
			fmt.Printf("%c", name[i]-32)

		}

	}
}
