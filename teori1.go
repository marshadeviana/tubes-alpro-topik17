package main

import "fmt"

const N int = 1000

func sudahTerurut(T [N]int, total int) bool {
    for i := 0; i < total-1; i++ {
        if T[i] >= T[i+1] {
            return false
        }
    }
    return true
}

func main() {
    var T [N]int
    var total int
    fmt.Scan(&total)
    for i := 0; i < total; i++ {
        fmt.Scan(&T[i])
    }
    fmt.Println(sudahTerurut(T, total))
}