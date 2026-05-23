package main

import "fmt"

func inpData(nama *string, status *string, angkutan *int) {
	fmt.Scan(nama, status, angkutan)
}

func outputData(nama string, status string, angkutan int) {
	fmt.Printf("| %-12s | %-15s | %-10d |\n", nama, status, angkutan)
}

func main() {
	var n, angkutan int
	var nama, status string

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		inpData(&nama, &status, &angkutan)
		if i == 1 {
			fmt.Println("--------------------------------------------------")
			fmt.Println("| DATA                                           |")
			fmt.Println("--------------------------------------------------")
		}
		outputData(nama, status, angkutan)
	}
}