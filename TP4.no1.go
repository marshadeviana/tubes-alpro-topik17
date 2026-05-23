package main

import "fmt"

func hitungLuasKelilingLingkaran(r float64, l *float64, k *float64) {
	*l = 3.14 * r * r
	*k = 2 * 3.14 * r
}

func hitungLuasKelilingPersegi(s float64, l *float64, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(lL, lP, kL, kP float64, toL *float64, toK *float64) {
	*toL = lL + lP
	*toK = kL + kP
}

func main() {

	var r, s float64
	var ll, lp float64
	var kl, kp float64
	var tl, tp float64
	
	fmt.Println("R S LL LP KL KP TL TP")

	for {
		fmt.Scan(&r, &s)

		if r == 0 && s == 0 {
			break
		}

		hitungLuasKelilingLingkaran(r, &ll, &kl)
		hitungLuasKelilingPersegi(s, &lp, &kp)
		hitungTotal(ll, lp, kl, kp, &tl, &tp)

		fmt.Printf("%.2f %.2f %.2f %.2f %.2f %.2f %.2f %.2f\n",
			r, s, ll, lp, kl, kp, tl, tp)
	}
}
