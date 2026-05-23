package main

import "fmt"

func numSeq(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func sumNumSeq(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += numSeq(i)
	}
	return total
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(sumNumSeq(n))
}