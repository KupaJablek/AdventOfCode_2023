package main

import (
	"fmt"
	"slices"
)


type star struct {
    r int
    c int
}

var stars []star
var eRows []int
var eCols []int
var data []string 

func parseData(d []string) {
    for i := 0; i < len(d); i++ {
        eCols = append(eCols, 0)
    }
    for i := 0; i < len(d[0]); i++ {
        eRows = append(eRows, 1)
    }

    for i := range d {
        rowEmpty := true
        for j := range d[i] {
            if d[i][j] == '#' {
                rowEmpty = false
                stars = append(stars, star{i, j})
                eCols[j] = 1
            }
        }
        if rowEmpty {
            eRows[i] = 0
        }
    }
}

func calcDist(s1, s2 star) (int, int) {
    dist := 0
    dist2 := 0

    cd := abs(s2.c - s1.c)
    rd := abs(s2.r - s1.r)
    dist += cd + rd
    dist2 += cd + rd

    cs := min(s1.c, s2.c) + 1
    rs := min(s1.r, s2.r) + 1

    ce := max(s1.c, s2.c)
    re := max(s1.r, s2.r)

    for i := cs; i < ce; i++ {
        if eCols[i] == 0 {
            dist += 1
            dist2 += 1000000 - 1
        } 
    }
    for i := rs; i < re; i++ {
        if eRows[i] == 0 {
            dist += 1
            dist2 += 1000000 - 1
        } 
    }

    return dist, dist2
}

func D11() {
    data, err := ParseFile("./11/input.txt")
    if err != nil {
        fmt.Println(err.Error())   
        return
    }

    parseData(data)

    total := 0
    total2 := 0
    var calcdStars []star
    for i := range stars {
        for j := range stars {
            if stars[j] == stars[i] || slices.Contains(calcdStars, stars[j]){
                continue
            }
            d, d2 := calcDist(stars[i], stars[j])
            total += d
            total2 += d2
        }
        calcdStars = append(calcdStars, stars[i])
    }

    fmt.Println("Part 1:", total)
    fmt.Println("Part 2:", total2)
}
