package main

import (
	"fmt"
	"reflect"
)

func d14siftRocks(m [][]string) [][]string {
    for i := range m {
        key := -1
        for k := range m[0] {
            r := m[k][i]
            if r == "." && key == -1 {
                key = k
            }
            if r == "O" && key != -1 {
                m[key][i] = "O"
                key++
                m[k][i] = "."
            } else if r == "#" {
                key = -1
            }
        }
    }
    return m
}

func d14rotate(m [][]string) [][]string {
    rotated := make([][]string, len(m))
    for i := range rotated {
        rotated[i] = make([]string, len(m[i]))
    }

    for i := range m {
        idx := len(m) - 1 - i
        for k := range m[i] {
            rotated[k][idx] = m[i][k]
        }
    }
    return rotated
}

func d14p2(m [][]string, cycles int) int {
    seen := [][][]string{}
    seen = append(seen, m)
    found := -1

    tm := m
    var c int
    for c = 1; c <= cycles; c++ {
        for x := 0; x < 4; x++ {
            tm = d14siftRocks(tm)
            tm = d14rotate(tm)
        }
        
        for i := range seen {
            if reflect.DeepEqual(seen[i], tm) {
                fmt.Println("FOUND IT")
                fmt.Println("\t", i)
                break
            }
        }

        if found != -1 {
            break
        }
        seen = append(seen, tm)
    }
    first := 0
    for i := range seen {
        if reflect.DeepEqual(seen[i], tm) {
            first = i
            fmt.Println("FOUND IT")
            fmt.Println("\t", i)
            break
        }
    }

    grid := seen[(cycles - first) % (c - first) + first]

    tot := 0
    for i := range grid {
        for k := range grid[i] {
            if grid[i][k] == "O" {
                tot += len(grid) - i
            }
        }
    }
    return tot
}

func d14() {
    d, err := ParseFile("14/intput.txt")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    m := make([][]string, len(d))
    for i := range d {
        m[i] = make([]string, len(d[i]))
        for k, val := range d[i] {
            m[i][k] = string(val)
        }
    }
    //cycles := 1000000000
    cycles := 3 
    p2 := d14p2(m, cycles)
    m = d14siftRocks(m)

    tot := 0
    for i := range m {
        for k := range m[i] {
            if m[i][k] == "O" {
                tot += len(m) - i
            }
        }
    }
    fmt.Println("Part 1:", tot)
    fmt.Println("Part 2:", p2)
}
