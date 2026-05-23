package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func bacaData(A *tabInt, n *int) {
	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX 
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int) {
	fmt.Print("Data Bilangan: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Println()
}

func sequentialSearch(A tabInt, n int, x int) bool {
	var ketemu bool = false
	var i int = 0
	for i < n && !ketemu {
		if A[i] == x {
			ketemu = true
		}
		i++
	}
	return ketemu
}

func frekuensiBilangan(A tabInt, n int, x int) int {
	var count int = 0
	for i := 0; i < n; i++ {
		if A[i] == x {
			count++
		}
	}
	return count
}

func main() {
	var data tabInt
	var nData, x1 int

	fmt.Scan(&x1)
	bacaData(&data, &nData)
	cetakData(data, nData)

	if sequentialSearch(data, nData, x1) {
		f := frekuensiBilangan(data, nData, x1)
		fmt.Printf("Hasil pencarian: Bilangan ditemukan. Terdapat %d bilangan %d.\n", f, x1)
	} else {
		fmt.Println("Hasil pencarian: Bilangan tidak ditemukan.")
	}
}