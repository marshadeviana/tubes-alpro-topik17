package main

import "fmt"

func main() {
	var x, n, i int
	var a int

	fmt.Scan(&x)
	fmt.Scan(&n)

	for n != 0 {
		i++
		if n == x && a == 0 {
			a = i
		}
		fmt.Scan(&n)
	}

	if a > 0 {
		fmt.Println(a)
	} else {
		fmt.Println("TIDAK ADA")
	}
}