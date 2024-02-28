package main

import (
	"fmt"
	"strings"
)

var alpha = [26]rune{
    'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k',
    'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
    'w', 'x', 'y', 'z',
}

func incrementPass(s string) string {
    pass := []rune(s)
    for i := len(pass) - 1; i >= 0; i-- {
        carry := false
        if pass[i] == 'z' {
            carry = true
        }
        pass[i] = alpha[(pass[i] - 'a' + 1) % 26]

        if !carry {
            break
        }
    }
    return string(pass)
}

func validPass(s string) bool {
    if strings.Contains(s, "i") || strings.Contains(s, "o") || strings.Contains(s, "l") {
        return false
    } 
    ls := len(s)
    l := 0
    run := false
    m := make(map[string]int)
    for r := 1; r < ls; r++ {
        for r - l > 2 {
           l++ 
        }
        if r - l == 2 && run == false {
            // check for run
            idx := s[r] - 'a'
            if idx >= 2 {
                if rune(s[r]) == alpha[idx] && rune(s[r - 1]) == alpha[idx - 1] && rune(s[r - 2]) == alpha[idx - 2]{
                    run = true
                }
            }
        }

        if s[r] == s[r - 1] {
            doub := string(s[r-1] + s[r])
            if _, ok := m[doub]; !ok {
                m[doub] = r
            }
        }
    }

    return len(m) >= 2 && run
}

func d11p1(d string) string {
    pass := d
    valid := false
    for !valid {
        pass = incrementPass(pass)
        valid = validPass(pass)
    }
    return pass
}

func d11() {
    d, err := ParseFile("11/input.txt")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    p1 := d11p1(d[0])
    fmt.Println("Part 1:", p1)
    fmt.Println("Part 2:", d11p1(p1))
}
