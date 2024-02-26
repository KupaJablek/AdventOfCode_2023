package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func encode(input string) string {
    return fmt.Sprintf("%x", md5.Sum([]byte(input)))
}

func d4find(d string, nz int) int {
    for i := 0; i < 10000000; i++ {
        res := encode(d + strconv.Itoa(i))

        valid := true
        for k := 0; k < nz; k++ {
            if res[k] != '0' {
                valid = false
                break
            } 
        }

        if valid {
            return i
        }
    }
    return 0
}

func d4() {
    d, err := ParseFile("04/input.txt") 
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("Part 1:", d4find(d[0], 5))
    fmt.Println("Part 2:", d4find(d[0], 6))
}
