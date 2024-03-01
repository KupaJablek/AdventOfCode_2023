package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Ingredient struct {
    cap, dur, fla, tex, cal int
}

func d15calc(ing []Ingredient, calCap int) int {
    spoons := 100
    maxCookie := 0
    for i := 0; i < spoons; i++ {
        for k := 0; k < spoons; k++ {
            for j := 0; j < spoons; j++ {
                c4 := spoons - i - k - j
                cap := ing[0].cap * i + ing[1].cap * k + ing[2].cap * j + ing[3].cap * c4
                dur := ing[0].dur * i + ing[1].dur * k + ing[2].dur * j + ing[3].dur * c4
                fla := ing[0].fla * i + ing[1].fla * k + ing[2].fla * j + ing[3].fla * c4
                tex := ing[0].tex * i + ing[1].tex * k + ing[2].tex * j + ing[3].tex * c4
                cals := ing[0].cal * i + ing[1].cal * k + ing[2].cal * j + ing[3].cal * c4

                if cap <= 0 || dur <= 0 || fla <= 0 || tex <= 0 {
                    continue
                }
                if calCap != 0 && cals != calCap {
                    continue
                }

                score := cap * dur * fla * tex
                if score > maxCookie {
                    maxCookie = score
                }
            }
        }
    }
    return maxCookie
}

func d15() {
    d, e := ParseFile("15/input.txt")
    if e != nil {
        fmt.Println(e.Error())
        return
    }

    ing:= make([]Ingredient, len(d))
    for i := range d {
        bits := strings.Split(d[i], ",") 
        cap, _ := strconv.Atoi(strings.Split(bits[0], " ")[2])
        dur, _ := strconv.Atoi(strings.Split(bits[1], " ")[2])
        fla, _ := strconv.Atoi(strings.Split(bits[2], " ")[2])
        tex, _ := strconv.Atoi(strings.Split(bits[3], " ")[2])
        cal, _ := strconv.Atoi(strings.Split(bits[4], " ")[2])
        ing[i] = Ingredient{cap, dur, fla, tex, cal}
    }

    fmt.Println("Part 1:", d15calc(ing, 0))
    fmt.Println("Part 2:", d15calc(ing, 500))
}
