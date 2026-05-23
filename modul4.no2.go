package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var computerNumber, yourNumber int
	var win bool = false
	var round int = 0
	
	fmt.Println("Start")
	for {
		round++
		fmt.Println("Round", round)
		numGenerator(&computerNumber)
		yourGuessing(&yourNumber)
		process(&yourNumber, &computerNumber, &win)
		
		if win || round > 4 {
			break
		}
	}
	conclusion(round, win)
	fmt.Println("End")
}

func yourGuessing(yN *int) {
	fmt.Print("Enter your guess: ")
	fmt.Scan(yN)
}

func numGenerator(cN *int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	*cN = r1.Intn(5)
}

func process(yN *int, cN *int, w *bool) {
	*w = *yN == *cN
	fmt.Printf("Your guessing: %d, computer number: %d, win: %t\n", *yN, *cN, *w)
}

func conclusion(r int, w bool) {
	if r <= 5 && w {
		fmt.Printf("You win in %d round\n", r)
	} else {
		fmt.Printf("Computer win in %d round\n", r)
	}
}
	