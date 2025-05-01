package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    dice := rand.Intn(6) + 1

    fmt.Println("🎲 You rolled:", dice)

    switch dice {
    case 1:
        fmt.Println("Move 1 step. Slow and steady!")
    case 2:
        fmt.Println("Move 2 steps. Getting started!")
    case 3:
        fmt.Println("Move 3 steps. Not bad!")
    case 4:
        fmt.Println("Move 4 steps. Good move!")
    case 5:
        fmt.Println("Move 5 steps. Almost there!")
    case 6:
        fmt.Println("Move 6 steps. 🎉 Bonus turn!")
    }
}
