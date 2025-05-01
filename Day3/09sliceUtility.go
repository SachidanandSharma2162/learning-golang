package main

import (
    "fmt"
    "sort"
)

func main() {
    var numbers []int
    var n int

    fmt.Println("Kitne numbers daalne hain?")
    fmt.Scan(&n)

    fmt.Println("Numbers daalo:")
    for i := 0; i < n; i++ {
        var temp int
        fmt.Scan(&temp)
        numbers = append(numbers, temp)
    }

    fmt.Println("\n📦 Slice:", numbers)
    fmt.Println("➕ Sum:", getSum(numbers))
    fmt.Println("🧹 Unique:", removeDuplicates(numbers))
    fmt.Println("🔥 Top 3:", getTopThree(numbers))
}

func getSum(slice []int) int {
    sum := 0
    for _, v := range slice {
        sum += v
    }
    return sum
}

func removeDuplicates(slice []int) []int {
    seen := make(map[int]bool)
    result := []int{}
    for _, v := range slice {
        if !seen[v] {
            seen[v] = true
            result = append(result, v)
        }
    }
    return result
}

func getTopThree(slice []int) []int {
    if len(slice) < 3 {
        return slice
    }
    temp := make([]int, len(slice))
    copy(temp, slice)
    sort.Sort(sort.Reverse(sort.IntSlice(temp)))
    return temp[:3]
}
