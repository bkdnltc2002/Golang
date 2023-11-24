// run  print hello 1->1000
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println("Hello ", i+1)
			fmt.Println("World")
		}()
	}
}
