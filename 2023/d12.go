package main

import (
	"fmt"
	"strings"
)

func getArrangements(line string) int {
    total := 0

    data := strings.Split(line, " ")
    s, groups := data[0], strings.Split(data[1], ",")

    fmt.Println(s)
    fmt.Println(groups)
    // . = good, # = not good, ? = unknown
    // groups of # are seperated by .

    return total
}

func d12() {
    fmt.Println("Day 12")
    //data, err := ParseFile("./12/input.txt")
    data, err := ParseFile("./12/test.txt")
    if err != nil {
        fmt.Println(err.Error())   
        return
    }

    total := 0
    for d := range data {
        total += getArrangements(data[d])
        fmt.Println(data[d])
    }

    fmt.Println("Part 1:", total)
}
