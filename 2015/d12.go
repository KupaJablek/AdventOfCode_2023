package main

import (
	"fmt"
	"strconv"
	"unicode"
    "encoding/json"
)

func d12p1(d []string) int {
    tot := 0
    for i := range d {
        num := 0
        neg := false
        for k := range(d[i]) {
            r := rune(d[i][k])
            if !unicode.IsNumber(r){
                if neg {
                    num *= -1 
                    neg = false
                }
                tot += num
                num = 0
                continue
            }

            if num == 0 && k - 1 >= 0 {
                if d[i][k - 1] == '-' {
                    neg = true
                }
            }
            res, _ := strconv.Atoi(string(r))
            num = 10 * num + res
        }

        if num > 0 {
            if neg {
                num *= -1 
            }
            tot += num
        }
    }
    return tot
}

func d12p2(j []string) int {
    sToParse := ""
    for r := range j {
        sToParse += j[r]
    }
    var result map[string]any
    json.Unmarshal([]byte(sToParse), &result)
    
    tot := 0
    for key, _ := range result {
        fmt.Println(key)
    }
    return tot
}

func d12() {
    data, err := ParseFile("12/input.txt")
    if err != nil {
        fmt.Println("input file not found")
        return
    }

    fmt.Println("Part 1:", d12p1(data))
    fmt.Println("Part 2:", d12p2(data))
}
