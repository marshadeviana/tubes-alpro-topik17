package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		fmt.Print(i)
		if i > 1 {
			fmt.Print(" x ")
		}
	}
	fmt.Println()
}