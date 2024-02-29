package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Deer struct {
    speed, fTime, rTime, points int
    resting bool
    tFly, tRested, dt int
}

func makeDeer(s []string) []Deer {
    deer := make([]Deer, len(s))
    for i := range s {
        var td Deer
        d := strings.Split(s[i], " ")
        speed, _ := strconv.Atoi(d[3])
        fTime, _ := strconv.Atoi(d[6])
        rTime, _ := strconv.Atoi(d[len(d) - 2])

        td.speed = speed
        td.fTime = fTime
        td.rTime = rTime
        deer[i] = td
    }
    return deer
} 

func runDeer(deer []Deer, time int) (int, int) {
    var maxD, maxP int 
    for i := 1; i <= time; i++ {
        for d, v := range deer {
            if v.resting {
                if v.tRested == v.rTime {
                    deer[d].resting = false
                    deer[d].tRested = 0
                } else {
                    deer[d].tRested++
                    continue
                }
            }

            deer[d].tFly++
            deer[d].dt += v.speed

            if deer[d].dt > deer[maxD].dt {
                maxD = d
            }

            if deer[d].tFly == v.fTime {
                deer[d].tFly = 0
                deer[d].resting = true
            }
        }

        for k, d := range deer {
            if d.dt == deer[maxD].dt {
                deer[k].points++
            }
            if deer[k].points > deer[maxP].points {
                maxP = k
            }
        }
    }
    return deer[maxD].dt, deer[maxP].points
}

func d14() {
    d, err := ParseFile("14/input.txt")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    deer := makeDeer(d)
    md, mp := runDeer(deer, 2503)
    fmt.Println("Part 1:", md)
    fmt.Println("Part 2:", mp)
}
