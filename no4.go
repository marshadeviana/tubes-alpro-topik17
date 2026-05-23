package main

import (
	"fmt"
	"math"
)

const g = 9.8

func konversiDerajatkeRadian(T float64) float64 {
	return T * (math.Pi / 180.0)
}

func waktu(V, T float64) float64 {
	rad := konversiDerajatkeRadian(T)
	return (V * math.Sin(rad)) / g
}

func jarak(V, T float64) float64 {
	rad := konversiDerajatkeRadian(T)
	return (math.Pow(V, 2) * math.Sin(2*rad)) / g
}

func ketinggian(V, T float64) float64 {
	rad := konversiDerajatkeRadian(T)
	return (math.Pow(V, 2) * math.Pow(math.Sin(rad), 2)) / (2 * g)
}

func main() {
	var V, T float64
	fmt.Scan(&V, &T)

	resWaktu := waktu(V, T)
	resJarak := jarak(V, T)
	resTinggi := ketinggian(V, T)

	fmt.Printf("%.2f %.2f %.2f\n", resWaktu, resJarak, resTinggi)
}