package main
import "fmt"

func reamur(c float64) float64 {
	return 4.0/5.0*c
}

func fahrenheit(c float64) float64 {
	return 9.0/5.0*c + 32
}

func kelvin(c float64) float64 {
	return c + 273
}

func main() {
	var awal, akhir, step float64
	var c float64

	fmt.Scan(&awal, &akhir, &step)

	fmt.Println("Celsius Reamur Fahrenheit Kelvin")

	for c = awal; c <= akhir; c = c + step {
		fmt.Printf("%.2f %.2f %.2f %.2f\n", c, reamur(c), fahrenheit(c), kelvin(c))
	}
}
