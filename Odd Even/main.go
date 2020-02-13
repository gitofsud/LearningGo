package main

import "fmt"

func main() {
	n := [11]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, x := range n {
		fmt.Print(x, " is ")
		if x%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
