package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func baca(A *tabInt, n *int) {
	var input int
	*n = 0
	for {
		fmt.Scan(&input)
		if input == 0 || *n >= NMAX {
			break
		}
		if input < 0 {
			input = -input
		}
		A[*n] = input
		*n++
	}
}

func cetak(A tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	res := 0
	for i := 0; i < n; i++ {
		res += A[i]
	}
	return res
}

func rataRata(A tabInt, n int) float64 {
	if n == 0 {
		return 0.0
	}
	return float64(jumlah(A, n)) / float64(n)
}

func main() {
	var data tabInt
	var nData int

	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), rataRata(data, nData))
}