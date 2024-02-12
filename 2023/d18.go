package main

import (
	"fmt"
	"strconv"
	"strings"
)

type point struct {
    x,y int
}

var points []point
var perimeter int

var points2 []point
var perimeter2 int

func dig(s []string) {
    curr := point{0, 0}
    points = append(points, curr)

    for i := range s {
        bits := strings.Split(s[i], " ")
        l, _ := strconv.Atoi(string(bits[1]))
        perimeter += l 

        p := curr

        switch string(bits[0]) {
        case "R":
            p.x += l
        case "L":
            p.x -= l
        case "D":
            p.y += l
        case "U":
            p.y -= l
        }
        points = append(points, p)
        curr = p
    }
}

func dig2(s []string) {
    curr := point{0, 0}
    points2 = append(points2, curr)
    p := curr

    for i := range s {
        bits := strings.Split(s[i], " ")[2]

        h, _ := strconv.ParseInt(bits[2:len(bits) - 2], 16, 64)
        l := int(h)

        perimeter2 += l

        switch string(bits[len(bits) - 2]) {
        case "0":
            p.x += l
        case "2":
            p.x -= l
        case "1":
            p.y += l
        case "3":
            p.y -= l
        }
        points2 = append(points2, p)
        curr = p
    }
}

func calc(p []point) int {
    var tot int
    for i := 1; i < len(p); i += 2 { 
        p1 := p[i - 1] 
        p2 := p[i]
        tot += (p2.x - p1.x) * (p2.y + p1.y)
    }

    return abs(int(tot) / 2) + int(perimeter2) / 2 + 1
}

func d18() {
    fmt.Println("Day 18")
    data, err := ParseFile("./18/input.txt")
    if err != nil {
        fmt.Println(err.Error())   
        return
    }

    dig(data)
    dig2(data)

    fmt.Println("Part #1:", calc(points))
    fmt.Println("Part #2:", calc(points2))
}
