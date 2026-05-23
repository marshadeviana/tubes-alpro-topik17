package main

import "fmt"

func printNumRecursive(n int) {
	if n > 0 {
		printNumRecursive(n - 1)
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	printNumRecursive(n)
	fmt.Println()
}