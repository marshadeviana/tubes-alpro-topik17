package main

import "fmt"

func main() {
	const NMAX = 10
	var A [NMAX]int
	var n, i, min int

	fmt.Scan(&n)

	if n > NMAX {
		n = NMAX
	}

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	min = 0

	for i = 1; i < n; i++ {
		if A[i] < A[min] {
			min = i
		}
	}

	fmt.Println(A[min], min)
}