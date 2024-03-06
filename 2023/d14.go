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

var tm [][]string
func cycle() {
    for x := 0; x < 4; x++ {
        tm = d14siftRocks(tm)
        tm = rotateArr(tm)
    }
}

func d14p2(cycles int) int {
    seen := make([][][]string, 150)
    
    seen[0] = make([][]string, len(tm))
    for i := range tm {
        seen[0][i] = tm[i]
    }

    found := -1
    var c int
    for c = 1; c <= cycles; c++ {
        fmt.Println()
        for j := range tm {
            fmt.Println(tm[j])
        }
        fmt.Println()
        for j := range tm {
            fmt.Println(seen[0][j])
        }
        cycle()
        fmt.Println()
        for k := range tm {
            fmt.Println(seen[0][k])
        }
        for i := range seen {
            if reflect.DeepEqual(tm, seen[i]) {
                found = i
                fmt.Println("THATS THE ONE RIGHT THERE")
                fmt.Println("\t", i)

                break
            }
        }
        if found != -1 {
            break
        }
        seen[c] = make([][]string, len(tm))
        for i := range tm {
            seen[c][i] = tm[i]
        }
    }

    fmt.Println("Length of seen:", len(seen))

    first := 0
    /*
    if reflect.DeepEqual(seen[117], seen[103]) {
        fmt.Println("I THGINK ITS WORKING")
        fmt.Println(seen[117])
    }
    */
    for i := range seen {
        if reflect.DeepEqual(tm, seen[i]) {
            first = i
            fmt.Println("found at:", i)
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
    d, err := ParseFile("14/test.txt")
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
    //cycles := 125 
    tm = m
    p2 := d14p2(10)
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
