package main

import (
	"fmt"
	"strconv"
)

func d10_look_and_say(d string, c int) int {
    for i := 0; i < c; i++ {
        newDigits := ""
        num := string(d[0])
        c := 0
        for k := 0; k < len(d); k++ {
            n := string(d[k])

            if num != n {
                newDigits += strconv.Itoa(c) + num
                c = 0
                num = n
            }
            c++
        }
        d = newDigits + "x"
    }
    return len(d) - 1
}

func d10() {
    data, err := ParseFile("10/input.txt")
    if err != nil {
        fmt.Println("input file not found")
        return
    }

    fmt.Println("Part 1:", d10_look_and_say(data[0] + "x", 40))
    fmt.Println("Part 2:", d10_look_and_say(data[0] + "x", 50))
}
