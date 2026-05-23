package main

import "fmt"

const NMAX int = 5

type tabInt struct {
	info [NMAX]int
	n    int
}

func bacaData(A *tabInt) {
	var temp int
	fmt.Scan(&temp)
	if A.n < NMAX {
		A.info[A.n] = temp
		A.n++
	}
}

func cetakData(A tabInt) {
	if A.n == 0 {
		fmt.Println("Array kosong")
	} else {
		for i := 0; i < A.n; i++ {
			fmt.Print(A.info[i], " ")
		}
		fmt.Println()
	}
}

func main() {
	var data tabInt
	data.n = 0

	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)
	bacaData(&data)

	cetakData(data)
}