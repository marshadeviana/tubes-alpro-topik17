package main

import "fmt"

const NMAX = 20

type tabInt [NMAX]int

func baca(A *tabInt, n *int) {
	var input int
	*n = 0
	for *n < NMAX {
		fmt.Scan(&input)
		if input <= 0 {
			break
		}
		A[*n] = input
		*n++
	}
}

func cetakElemen(A tabInt, n int) {
	fmt.Print("Elemen array: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func maksimum(A tabInt, n int) int {
	max := A[0]
	for i := 1; i < n; i++ {
		if A[i] > max {
			max = A[i]
		}
	}
	return max
}

func minimum(A tabInt, n int) int {
	min := A[0]
	for i := 1; i < n; i++ {
		if A[i] < min {
			min = A[i]
		}
	}
	return min
}

func cetakInfo(A tabInt, n int) {
	if n > 0 {
		fmt.Printf("Nilai maksimum: %d\n", maksimum(A, n))
		fmt.Printf("Nilai minimum: %d\n", minimum(A, n))
		fmt.Printf("Banyak elemen: %d\n", n)
	}
}

func main() {
	var data tabInt
	var n int

	baca(&data, &n)
	if n > 0 {
		cetakElemen(data, n)
		cetakInfo(data, n)
	}
}