package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getHash(s string) int {
    var total int
    for j := 0; j < len(s); j++ {
        total = ((total + int(s[j])) * 17) % 256
    }
    return total
}

type lensePair struct {
    label string
    flen int
}

func calculateP2(steps []string) int {
    var total int
    boxes := make(map[int][]lensePair)

    for i := 0; i < len(steps); i++ {
        l := steps[i]
        if l[len(l) - 1] == '-' {
            label := l[:len(l) - 1]
            index := getHash(label)
            for j := 0; j < len(boxes[index]); j++ {
                if boxes[index][j].label == label {
                    boxes[index] = append(boxes[index][:j],boxes[index][j+1:]...)
                    break
                } 
            }
        } else {
            label := l[:len(l) - 2]
            index := getHash(label)
            flen, _ := strconv.Atoi(string(l[len(l) - 1]))

            box := false
            for j := 0; j < len(boxes[index]); j++ {
                if boxes[index][j].label == label {
                    boxes[index][j].flen = flen
                    box = true
                    break
                } 
            }
            if !box {
                boxes[index] = append(boxes[index], lensePair{label, flen})
            }
        }
    }

    for index, items := range boxes {
        for l := 0; l < len(items); l++ {
            total += (index + 1) * (l + 1) * (items[l].flen)
        }
    }

    return total
}

func d15() {
    data, err := ParseFile("./15/input.txt")
    if err != nil {
        fmt.Println(err.Error())   
        return
    }

    steps := strings.Split(data[0], ",")

    p1 := 0
    for i := 0; i < len(steps); i++ {
        p1 += getHash(steps[i])
    }

    fmt.Println("Part 1:", p1)
    fmt.Println("Part 2:", calculateP2(steps))
}
