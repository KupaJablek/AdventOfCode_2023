package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part struct {
    x, m, a, s int
}

func testConds(c string, p Part) string {
    opts := strings.Split(c, ",")

    for i := range opts {
        stat := string(opts[i][0])
        loc := strings.Split(opts[i], ":")
        if len(loc) == 1 {
            return loc[0]
        }

        val, _:= strconv.Atoi(loc[0][2:])
        check := string(opts[i][1])

        switch stat {
            case "x": 
            if check == "<" && p.x < val {
                return loc[1]
            }
            if  check == ">" && p.x > val {
                return loc[1]
            }
            case "m": 
            if check == "<" && p.m < val {
                return loc[1]
            }
            if  check == ">" && p.m > val {
                return loc[1]
            }
            case "a": 
            if check == "<" && p.a < val {
                return loc[1]
            }
            if  check == ">" && p.a > val {
                return loc[1]
            }
            case "s": 
            if check == "<" && p.s < val {
                return loc[1]
            }
            if  check == ">" && p.s > val {
                return loc[1]
            }
        }
    }
    return ""
}

func doSteps(m map[string]string, p Part) bool {
    cStep := "in"
    //fmt.Println(p)
    //fmt.Print("\t")
    for {
        //fmt.Print(cStep, " -> ")
        if val, ok := m[cStep]; ok {
            conds := testConds(val, p)
            if conds == "A" {
                //fmt.Print("A\n")
                return true
            }
            if conds == "R" {
                //fmt.Print("R\n")
                return false
            }
            cStep = conds
        } else {
            fmt.Println("KEY NOT FOUND!!!!", cStep)
            break
        }
    }
    return false
}

func d19() {
    d, err := ParseFile("19/input.txt")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    steps := make(map[string]string)
    var line int
    for line = 0; d[line] != ""; line++ {
        b := strings.Split(d[line], "{")
        steps[string(b[0])] = b[1][: len(b[1]) - 1] 
    } 

    line++ // hit blank

    parts := make([]Part, 0)
    for line < len(d) {
        bits := d[line][1:len(d[line]) - 1]
        counts := strings.Split(bits, ",")
        x, _ := strconv.Atoi(string(counts[0][2:]))
        m, _ := strconv.Atoi(string(counts[1][2:]))
        a, _ := strconv.Atoi(string(counts[2][2:]))
        s, _ := strconv.Atoi(string(counts[3][2:]))
        tmp := Part{x, m, a, s}
        parts = append(parts, tmp)
        line++
    } 

    tot := 0
    for p := range parts {
        if doSteps(steps, parts[p]) {
            tot += parts[p].x + parts[p].m + parts[p].a + parts[p].s
        }
    }
    fmt.Println("Part 1:", tot)


    // all combinations between 1-4000 for all part values

}
