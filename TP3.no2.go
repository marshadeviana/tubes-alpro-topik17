package main
import "fmt"

func lowToUpper(k byte) byte {
	return k - 32
}

func main() {
	var h byte

	fmt.Scanf("%c", &h)

	fmt.Printf("%c\n", lowToUpper(h))
}
