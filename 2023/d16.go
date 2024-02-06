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

var maze []string
var count int
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
    // step another space forward
    switch w.direction {
    case Up:
        //fmt.Println("going up")
        w.l.y--
    case Down:
        //fmt.Println("going down")
        w.l.y++
    case Right:
        //fmt.Println("going right")
        w.l.x++
    case Left:
        //fmt.Println("going left")
        w.l.x--
    }
    return w
}

func walk(w work) []work {
    out := []work{}

    for validIdx(w.l) {
        count++
        // check for direction change
        c := string(maze[w.l.y][w.l.x])
        //fmt.Println(c)
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
                    return []work{}
                }
                splits = append(splits, w.l)

                n1 := work{w.l, Up}
                n2 := work{w.l, Down}
                n1 = increment(n1)
                n2 = increment(n2)

                out = append(out, n1, n2)
                return out
            default: // nothing
            }
        case "-":
            switch w.direction {
            case Up, Down:
                if slices.Contains(splits, w.l) {
                    return []work{}
                }
                splits = append(splits, w.l)

                n1 := work{w.l, Left}
                n2 := work{w.l, Right}
                n1 = increment(n1)
                n2 = increment(n2)

                out = append(out, n1, n2)
                return out
            default: // nothing
            }
        default: // "." -> do nothing 
        }
        w = increment(w)
    }
    return out
}

func d16() {
    //data, err := ParseFile("./16/input.txt")
    data, err := ParseFile("./16/test.txt")
    if err != nil {
        fmt.Println(err.Error())   
        return
    }

    maze = data
    for i := range data {
        fmt.Println(maze[i])
    }

    tasks := []work{work{loc{0,0}, Right}}

    for len(tasks) > 0 {
        t := tasks[0]
        tasks = tasks[1:]
        r := walk(t)
        if len(r) > 0 {
            //fmt.Println(r)
            for i := range r {
                tasks = append(tasks, r[i])
            }
        }
    }

    fmt.Println("Part #1:", count)
}
