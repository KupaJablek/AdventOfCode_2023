package main

import (
	"fmt"
	"strconv"
	"strings"
)

func d2p1(data []string) int {
    tot := 0
    for i := range data {
        d := strings.Split(data[i], "x")
        n := make([]int, 3)
        for i := range d {
            n[i], _ = strconv.Atoi(d[i])
        }

        sa := (2 * n[0] * n[1]) + (2 * n[1] * n[2]) + (2 * n[2] * n[0])

        sm := n[0] * n[1]
        for k := range n {
            s := n[k] * n[(k + 1 ) % 3]
            if s < sm { sm = s }
        }
        tot += sa + sm
    }
    return tot
}

func d2p2(data []string) int {
    tot := 0
    for i := range data {
        d := strings.Split(data[i], "x")
        n := make([]int, 3)
        for i := range d {
            n[i], _ = strconv.Atoi(d[i])
        }
        
        bow := n[0] * n[1] * n[2]

        ribbon := n[0] * 2 + n[1] * 2
        for k := range n {
            p := n[k] * 2 + n[(k + 1 ) % 3] * 2
            if p < ribbon {
                ribbon = p
            }
        }
        tot += bow + ribbon
    }
    return tot
}

func d2() {
    data, err := ParseFile("02/input.txt")
    if err != nil {
        fmt.Println("input file not found")
        return
    }
    fmt.Println("Part 1: ", d2p1(data))
    fmt.Println("Part 2: ", d2p2(data))
}
