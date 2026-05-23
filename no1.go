package main

import "fmt"

func vowel_uppercase(kar byte) bool {
	return kar == 'A' || kar == 'I' || kar == 'U' || kar == 'E' || kar == 'O'
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
		fmt.Scanf("%c", &abjad) 
	}

	fmt.Printf("jumlah vokal uppercase adalah %d\n", jumlah)
}
