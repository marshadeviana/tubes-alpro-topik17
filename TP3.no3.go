package main
import "fmt"

func adaBilanganM(n, m int) string {

	for n > 0 {
		if n%10 == m {
			return "YA"
		}
		n = n / 10
	}

	return "TIDAK"
}

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	fmt.Println(adaBilanganM(n, m))
}
