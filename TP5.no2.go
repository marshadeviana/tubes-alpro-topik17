package main

import "fmt"

func f(a, b int) int {
	if b == 0 {
		return a
	}
	return f(b, a%b)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Println(f(x, y))
}