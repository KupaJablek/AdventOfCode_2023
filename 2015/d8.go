package main

import (
	"encoding/hex"
	"fmt"
)

func d8p1(d []string) int {
    var totC, totM int
    for _, val := range d {
        temp := ""
        for _ = range val {
            totC++
        }
        val = val[1:len(val) - 1]
        for i := 0; i < len(val); i++ {
            if val[i] == '\\' {
                if i + 1 < len(val) {
                    switch val[i + 1] {
                    case 'x':
                        h := val[i:i+4]
                        bs, _ := hex.DecodeString(h[2:])
                        temp += string(bs)
                        i += 3
                    case '\\':
                        temp += "\\"
                        i++
                    case '"':
                        temp += "\""
                        i++
                    }
                }
            } else {
                temp += string(val[i])
            }
        }
        totM += len(temp)
    }
    return totC - totM
}

func d8p2(d []string) int {
    var totC, totM int
    for _, val := range d {
        totM += 2
        for i := range val {
            totC++
            totM++
            if val[i] == '"' || val[i] == '\\' {
                totM++
            }
        }
    }
    return totM - totC
}

func d8() {
    data, err := ParseFile("08/input.txt")
    if err != nil {
        fmt.Println("input file not found")
        return
    }
    fmt.Println("Part 1: ", d8p1(data))
    fmt.Println("Part 2: ", d8p2(data))
}
