package main

import (
    "fmt"
)

func main() {
    var marks int
    fmt.Print("Enter your marks: ")
    fmt.Scanf("%d",&marks)
    // fmt.Scan(&marks)

    if marks >= 90 {
        fmt.Println("Grade: A")
    } else if marks >= 75 {
        fmt.Println("Grade: B")
    } else if marks >= 60 {
        fmt.Println("Grade: C")
    } else if marks >= 40 {
        fmt.Println("Grade: D")
    } else {
        fmt.Println("Grade: F (Fail)")
    }
}
