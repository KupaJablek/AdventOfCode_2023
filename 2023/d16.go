package main

import (
	"fmt"
	"slices"
)

const (
    Up int = iota
    Right
    Down
    Left
)

type loc struct {
    x int
    y int
}

type work struct {
    l loc 
    direction int 
}

var seen []loc
var maze []string
var splits []loc

func validIdx(l loc) bool {
    if l.x > len(maze) - 1 || l.x < 0 {
        return false
    }
    if l.y > len(maze[0]) - 1 || l.y < 0 {
        return false
    }
    return true
}

func increment(w work) work {
    switch w.direction {
    case Up:
        w.l.y--
    case Down:
        w.l.y++
    case Right:
        w.l.x++
    case Left:
        w.l.x--
    }
    return w
}

func walk(w work) ([]work, int) {
    out := []work{}
    count := 0

    for validIdx(w.l) {
        if !slices.Contains(seen, w.l) {
            count++
            seen = append(seen, w.l)
        }
        c := string(maze[w.l.y][w.l.x])
        switch c {
        case "\\":
            switch w.direction {
            case Up:
                w.direction = Left
            case Down:
                w.direction = Right
            case Left:
                w.direction = Up
            case Right:
                w.direction = Down
            }
        case "/":
            switch w.direction {
            case Down:
                w.direction = Left
            case Up:
                w.direction = Right
            case Right:
                w.direction = Up
            case Left:
                w.direction = Down
            }
        case "|":
            switch w.direction {
            case Left, Right:
                if slices.Contains(splits, w.l) {
                    return []work{}, count
                }
                splits = append(splits, w.l)

                n1 := work{w.l, Up}
                n2 := work{w.l, Down}
                n1 = increment(n1)
                n2 = increment(n2)

                out = append(out, n1, n2)
                return out, count
            default: // nothing
            }
        case "-":
            switch w.direction {
            case Up, Down:
                if slices.Contains(splits, w.l) {
                    return []work{}, count
                }
                splits = append(splits, w.l)

                n1 := work{w.l, Left}
                n2 := work{w.l, Right}
                n1 = increment(n1)
                n2 = increment(n2)

                out = append(out, n1, n2)
                return out, count
            default: // nothing
            }
        default: // "." -> do nothing 
        }
        w = increment(w)
    }
    return out, count
}

func p1(l loc, d int) int {
    total := 0

    tasks := []work{work{l, d}}
    for len(tasks) > 0 {
        t := tasks[0]
        tasks = tasks[1:]
        r, c := walk(t)
        total += c
        if len(r) > 0 {
            for i := range r {
                tasks = append(tasks, r[i])
            }
        }
    }
    return total
}

func clearGlobals() {
    seen = []loc{}
    splits = []loc{}
}

func p2() int {
    tot := 0
    for i := 0; i < len(maze); i++ {
        c := p1(loc{0, i}, Right)
        if c > tot { tot = c}
        clearGlobals()
    }
    for i := 0; i < len(maze); i++ {
        c := p1(loc{len(maze) - 1, i}, Left)
        if c > tot { tot = c}
        clearGlobals()
    }
    for i := range maze[0] {
        c := p1(loc{i, len(maze[0]) - 1}, Up)
        if c > tot { tot = c}
        clearGlobals()
    }
    for i := 0; i < len(maze[0]); i++ {
        c := p1(loc{i, 0}, Down)
        if c > tot { tot = c}
        clearGlobals()
    }
    return tot
}

func d16() {
    data, err := ParseFile("./16/input.txt")
    if err != nil {
        fmt.Println(err.Error())   
        return
    }
    maze = data

    ret := p1(loc{0,0}, Right)
    fmt.Println("Part #1:", ret)
    clearGlobals()

    // P2
    fmt.Println("Part #2:", p2())
}
