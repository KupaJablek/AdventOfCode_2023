package main

import "fmt"

func findHouse(n int) int {
    houses := make([]int, n)

    for i := 1; i <= n / 10; i++ {
        for k := i; k <= n / 10; k += i {
            houses[k] += i * 10
        }
    }

    for i, h := range houses {
        if h >= n {
            return i
        }
    }

    return -1
}

func findHouseP2(n int) int {
    houses := make([]int, n)

    for i := 1; i <= n / 10; i++ {
        for k := i; k <= n / 10 && k / 50 != i; k += i {
            houses[k] += i * 11
        }
    }

    for i, h := range houses {
        if h >= n {
            return i
        }
    }

    return -1
}

func d20() {
    input := 33100000

    fmt.Println("Part 1:", findHouse(input))
    fmt.Println("Part 2:", findHouseP2(input))
}
