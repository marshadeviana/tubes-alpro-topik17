package main

import "fmt"

func fungsi_f(x, y, z float64) float64 {
	return (2 * x) / (x + y + z)
}

func fungsi_g(x, y float64) float64 {
	return y 
}

func main() {
	var b1, b2, b3 float64

	fmt.Scan(&b1, &b2, &b3)

	hasil_f := fungsi_f(b1, b2, b3)
	hasil_g := fungsi_g(b1, b2)

	fmt.Printf("%v %v\n", hasil_f, hasil_g)
}