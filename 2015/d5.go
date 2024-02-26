package main

import "fmt"

func isVowel(r rune) bool {
    if r == 'a' || r == 'o' || r == 'e' || r == 'i' || r == 'u' {
        return true
    }
    return false
}

func naughtyCombo(s string) bool {
    switch s {
    case "ab", "cd", "pq", "xy":
        return true
    }
    return false
}

func d5p1(d[]string) int {
    tot := 0
    for i := range d{
        double, nc := false, false
        vowels := 0
        if isVowel(rune(d[i][0])) {
            vowels++
        }

        for k := 1; k < len(d[i]); k++ {
            combo := string(d[i][k-1]) + string(d[i][k])
            if naughtyCombo(combo) {
                nc = true 
            }
            if d[i][k] == d[i][k - 1] {
                double = true
            }
            if isVowel(rune(d[i][k])) {
                vowels++
            }
        }
        if vowels < 3 || !double || nc {
            continue
        }
        tot++
    }
    return tot
}

func d5p2(d []string) int {
    tot := 0
    for i := range d {
        pair, olr := false, false // one letter repeat ie: xyx
        m := make(map[string]int, len(d[i]))

        var r, l int
        for r = 1; r < len(d[i]); r++ {
            combo := string(d[i][r-1]) + string(d[i][r])
            if val, ok := m[combo]; ok {
                if r - val >= 2 {
                    pair = true
                } 
            } else {
                m[combo] = r
            }

            if r - l > 2 {
                l++
            }
            
            if r - l == 2 && d[i][l] == d[i][r] {
                olr = true
            }
        }

        if pair && olr {
            tot++
        }
    }
    return tot
}

func d5() {
    data, err := ParseFile("05/input.txt")
    if err != nil {
        fmt.Println("input file not found")
        return
    }

    fmt.Println("Part 1:", d5p1(data))
    fmt.Println("Part 2:", d5p2(data))
}
