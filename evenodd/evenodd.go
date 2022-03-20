package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, no := range x {
		if no%2 == 0 {
			fmt.Println(no, "is Even")
		} else {
			fmt.Println(no, "is Odd")
		}
	}

}
