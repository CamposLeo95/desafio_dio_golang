package main

import "fmt"

func main() {
	numeroMax := 100

	for i := 0; i <= numeroMax; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("pin pan!")
		} else if i%3 == 0 {
			fmt.Println("pin")
		} else if i%5 == 0 {
			fmt.Println("pan!")
		} else {
			fmt.Println(i)
		}
	}

}
