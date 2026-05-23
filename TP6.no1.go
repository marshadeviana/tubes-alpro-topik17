package main

import "fmt"

func main() {
	const NMAX int = 10
	var A [NMAX]string
	var i, n int

	fmt.Scan(&n)

	if n > NMAX {
		n = NMAX
	}

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	for i = 0; i < n; i++ {
		fmt.Println(A[i])
	}
}