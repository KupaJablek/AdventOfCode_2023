package main

import (
	"fmt"
	"strconv"
	"strings"
)

func d6() {
    data, err := ParseFile("06/input.txt")
    if err != nil {
        fmt.Println("input file not found")
        return
    }

    lights := make([][]bool, 1000)
    lights2 := make([][]int, 1000)
    for i := range lights {
        lights[i] = make([]bool, 1000)
        lights2[i] = make([]int, 1000)
    }

    for i := range data {
        bits := strings.Split(data[i], " ")
        state, offset := 0, 0
        if len(bits) == 5 {
            offset = 1
            if bits[1] == "on" {
                state = 1
            } else {
                state = 2
            }
        }

        start := strings.Split(bits[1 + offset], ",")
        end := strings.Split(bits[3 + offset], ",")

        s1, _ := strconv.Atoi(start[0])
        e1, _ := strconv.Atoi(end[0])
        s2, _ := strconv.Atoi(start[1])
        e2, _ := strconv.Atoi(end[1])
        for i := s1; i <= e1; i++ {
            for k := s2; k <= e2; k++ {
                switch state {
                case 0:
                    lights[i][k] = !lights[i][k]
                    lights2[i][k] += 2 
                case 1:
                    lights[i][k] = true 
                    lights2[i][k]++ 
                case 2:
                    lights[i][k] = false 
                    if lights2[i][k] > 0 {
                        lights2[i][k]-- 
                    }
                }
            }
        }
    }

    tot := 0
    for i := range lights {
        for k := range lights[0] {
            if lights[i][k] {
                tot++
            }
        }
    }

    tot2 := 0
    for i := range lights2 {
        for k := range lights2[0] {
            tot2 += lights2[i][k]
        }
    }
    fmt.Println("Part 1:", tot)
    fmt.Println("Part 2:", tot2)
}
