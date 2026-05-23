package main

import "fmt"

func main () {
var pilihan int
	
	for {
		tampil_menu()
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Println("Hello")
		case 2:
			fmt.Println("Hi")
		case 3:
			fmt.Println("Good Morning")
		case 4:
			fmt.Println("Good Night")
		
			
		}
		if pilihan == 5 { 
		fmt.Println("end")
		break }
	}
}

func tampil_menu() {
	fmt.Println("--------------------------")
	fmt.Println("        M E N U           ")
	fmt.Println("--------------------------")
	fmt.Println("1. Hello")
	fmt.Println("2. Hi")
	fmt.Println("3. Good Morning")
	fmt.Println("4. Good Night")
	fmt.Println("5. Exit")
	fmt.Println("--------------------------")
	fmt.Print("Your Choice (1/2/3/4/5)? ")
}