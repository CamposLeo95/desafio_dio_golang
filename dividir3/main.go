package main

import "fmt"

func main() {
	numeroMax := 100

	for i := 0; i < numeroMax; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
