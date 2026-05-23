package main
import "fmt"

funcv vowel_uppercase(kar byte) bool {
	var VPC bool
	VPC = kar == 'A' || kar == '0' || kar == 'I' || kar == 'E' || kar == 'U'
	return VPC
}
func main() {
	var abjad byte
	var jumlah int
	
	fmt.Scanf("%c", &abjad)
	jumlah = 0
	
	for abjad != '.' {
		if vowel_uppercase(abjad) {
			jumlah = jumlah + 1
			
		}
		fmt.Scanf("%c", abjad)
	}
	fmt.Printf(jumlah vokal uppercase adalah %d\n", jumlah)
}