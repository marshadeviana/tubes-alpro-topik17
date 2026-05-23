package main

import "fmt"

const N int = 1000

func countThis(T [N]int, total, ini int) int {
    jumlah := 0
    for i := 0; i < total; i++ {
        if T[i] == ini {
            jumlah++
        } else if T[i] > ini {
            break
        }
    }
    return jumlah
}

func main() {
    var T [N]int
    var total, ini int
    fmt.Scan(&total)
    for i := 0; i < total; i++ {
        fmt.Scan(&T[i])
    }
    fmt.Scan(&ini)
    fmt.Println(countThis(T, total, ini))
}