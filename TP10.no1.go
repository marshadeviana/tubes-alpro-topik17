package main

import "fmt"

func main() {
	const NMAX = 10
	var A [NMAX]int
	var n, i, max int

	fmt.Scan(&n)

	if n > NMAX {
		n = NMAX
	}

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	max = 0

	for i = 1; i < n; i++ {
		if A[i] > A[max] {
			max = i
		}
	}

	fmt.Println(A[max], max)
}