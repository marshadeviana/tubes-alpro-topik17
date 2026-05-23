package main

import "fmt"

func main() {
	var member bool
	var a, b, c, d, e int
	var persenDiskon, persenCashback float64

	fmt.Scan(&member, &a, &b, &c, &d, &e)

	if ganjil(a, b, c, d, e) {
		persenDiskon = diskon(member, a, c, e)
	} else if genap(a, b, c, d, e) {
		persenCashback = cashback(member, b, d)
	} else {
		persenDiskon = diskon(member, a, c, e)
		persenCashback = cashback(member, b, d)
	}

	fmt.Printf("%.2f %.2f\n", persenDiskon, persenCashback)
}

func ganjil(a, b, c, d, e int) bool {
	return (a%2 != 0) && (b%2 != 0) && (c%2 != 0) && (d%2 != 0)
}

func genap(a, b, c, d, e int) bool {
	return (a%2 == 0) && (b%2 == 0) && (c%2 == 0) && (d%2 == 0)
}

func diskon(member bool, a, c, e int) float64 {
	var nilai float64

	nilai = 1.5 * float64(a+c+e)

	if member {
		nilai = nilai + nilai*0.25
	}

	if nilai > 50 {
		nilai = 50
	}

	return nilai
}

func cashback(member bool, b, d int) float64 {
	var nilai float64

	nilai = 2.5 * float64(b+d)

	if member {
		nilai = nilai + nilai*0.15
	}

	if nilai > 40 {
		nilai = 40
	}

	return nilai
}