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

func rotateArr(m [][]string) [][]string {
    // transpose
    for i := range m{
        for k := i; k < len(m[0]); k++ {
            m[i][k], m[k][i] = m[k][i], m[i][k]
        }
    }

    // reverse each row
    for i := range m{
        j, k := 0, len(m[0]) - 1
        for j < len(m[0]) / 2 {
            m[i][j], m[i][k - j] = m[i][k - j], m[i][j]
            j++
        }
    }
    return m
}

func d14p2(tm [][]string, cycles int) int {
    seen := make([][][]string, 150)
    seen[0] = make([][]string, len(tm))
    for i := range tm {
        temp := make([]string, len(tm[i]))
        for k := range tm[i] {
            temp[k] = tm[i][k]
        }
        seen[0][i] = temp
    }

    var c int
    for {
        c++
        found := -1
        for x := 0; x < 4; x++ {
            tm = d14siftRocks(tm)
            tm = rotateArr(tm)
        }
        for i := 0; i < c; i++ {
            if reflect.DeepEqual(tm, seen[i]) {
                found = i
                break
            }
        }
        if found != -1 {
            break
        }
        seen[c] = make([][]string, len(tm))
        for i := range tm {
            temp := make([]string, len(tm[i]))
            for k := range tm[i] {
                temp[k] = tm[i][k]
            }
            seen[c][i] = temp
        }
    }

    var first int
    for i := range seen {
        if reflect.DeepEqual(tm, seen[i]) {
            first = i
            break
        }
    }

    grid := seen[(cycles - first) % (c - first) + first]
    return d14Tot(grid)
}

func d14Tot(grid [][]string) int {
    var tot int
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
    d, err := ParseFile("14/input.txt")
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

    p2 := d14p2(m, 1000000000)
    m = d14siftRocks(m)
    fmt.Println("Part 1:", d14Tot(m))
    fmt.Println("Part 2:", p2)
}
