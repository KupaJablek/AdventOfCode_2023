package main

import (
	"fmt"
	"slices"
)

type house struct {
    x, y int
}


func d3p1(d []string) int {
    seen := make([]house, 0)
    santa := house{0,0}

    seen = append(seen, santa)

    for i := range d {
        for k := range d[i] {
            move(&santa, rune(d[i][k]))
            if !slices.Contains(seen, santa) {
                seen = append(seen, santa)
            }
        }
    }
    return len(seen)
}
func d3p2(d []string) int {
    seen := make([]house, 0)
    santa := house{0,0}
    robo := house{0,0}
    seen = append(seen, santa)

    s := true
    for i := range d {
        for k := range d[i] {
            m := rune(d[i][k])
            if s {
                move(&santa, m)
                if !slices.Contains(seen, santa) {
                    seen = append(seen, santa)
                }
            } else {
                move(&robo, m)
                if !slices.Contains(seen, robo) {
                    seen = append(seen, robo)
                }
            }
            s = !s
        }
    }
    return len(seen)
}

func move(h *house, m rune) {
    switch m {
    case '>':
        h.x++
    case '<':
        h.x--
    case '^':
        h.y++
    case 'v':
        h.y--
    }
}

func d3() {
    data, err := ParseFile("03/input.txt")
    if err != nil {
        fmt.Println("input file not found")
        return
    }
    fmt.Println("Part 1: ", d3p1(data))
    fmt.Println("Part 2: ", d3p2(data))
}
