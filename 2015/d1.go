package main

import "fmt"

func d1p1(s []string) int {
    tot := 0
    for i := range s {
        for k := range s[i] {
            if s[i][k] == '(' {
                tot++
            } else {
                tot--
            }
        }
    }
    return tot
}


func d1p2(s []string) int {
    tot := 0
    pos := 0
    for i := range s {
        for k := range s[i] {
            if s[i][k] == '(' {
                tot++
            } else {
                tot--
            }

            pos++

            if tot == -1 {
                return pos
            }

        }
    }
    return tot
}

func d1() {
    data, err := ParseFile("01/input.txt")
    if err != nil {
        fmt.Println("input file not found")
        return
    }

    fmt.Println("Part 1: ", d1p1(data))
    fmt.Println("Part 2: ", d1p2(data))
}
