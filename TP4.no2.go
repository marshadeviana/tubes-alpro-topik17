package main

import "fmt"

func hitungMenang(g, k int, jm *int) {
	if g > k {
		*jm = *jm + 1
	}
}

func hitungDraw(g, k int, jd *int) {
	if g == k {
		*jd = *jd + 1
	}
}

func hitungKalah(g, k int, jk *int) {
	if g < k {
		*jk = *jk + 1
	}
}

func hitungJumGolKebobolanSelisih(g, k int, jg *int, jk *int, jsg *int) {
	*jg = *jg + g
	*jk = *jk + k
	*jsg = *jg - *jk
}

func hitungJumPoint(jm, jd int, jp *int) {
	*jp = jm*3 + jd*1
}

func main() {

	var N int
	var g, k int

	var jm, jd, jk int
	var jg, jkb, jsg int
	var jp int

	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&g, &k)

		hitungMenang(g, k, &jm)
		hitungDraw(g, k, &jd)
		hitungKalah(g, k, &jk)

		hitungJumGolKebobolanSelisih(g, k, &jg, &jkb, &jsg)
	}

	hitungJumPoint(jm, jd, &jp)

	fmt.Println(N, jm, jd, jk, jg, jkb, jsg, jp)
}
