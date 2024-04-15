package main

import "fmt"

// weapons
// dmg = index + 4
var w = []int{8, 10, 25, 40, 74}

/*
armour
Armour = index
*/
var a = []int{0, 13, 31, 53, 75, 102}

/*
rings
can buy 0-2
dmg = idx <= 2
*/
var r = []int{25, 50, 100, 20, 40, 80}

func d21Boss(dmg, ac int) bool {
    boss, hero := 104, 100

    for hero > 0 {
        // hero
        hh := dmg - 1
        if hh <= 0 {
            hh = 1
        }

        boss -= hh
        if boss <= 0 {
            return true
        }

        // boss
        bossDmg := 8 - ac
        if bossDmg <= 0 {
            bossDmg = 1
        }

        hero -= bossDmg
    }

    return false
}

func d21() {
    lowestCost := 100000
    highestCost := 0

    // no rings
    for i, ci := range w {
        for k, ck := range a {
            runCost := ci + ck

            armor := k
            dmg := i + 4

            if d21Boss(dmg, armor) {
                if runCost < lowestCost {
                    lowestCost = runCost
                }
            } else {
                if runCost > highestCost {
                    highestCost = runCost
                }
            }
        }
    }
    
    // 1 ring
    for i, ci := range w {
        for k, ck := range a {
            for j, jr := range r {
                runCost := ci + ck + jr

                armor := k
                dmg := i + 4

                if j <= 2 {
                    dmg += j + 1
                } else {
                    armor += j - 2
                }

                if d21Boss(dmg, armor) {
                    if runCost < lowestCost {
                        lowestCost = runCost
                    }
                } else {
                    if runCost > highestCost {
                        highestCost = runCost
                    }
                }
            }
        }
    }

    // 2 rings
    for i, ci := range w {
        for k, ck := range a {

            for j, jr := range r {
                for j2, jr2 := range r {
                    if j == j2 {
                        continue
                    }
                    runCost := ci + ck + jr + jr2

                    armor := k
                    dmg := i + 4

                    if j <= 2 {
                        dmg += j + 1
                    } else {
                        armor += j - 2
                    }

                    if j2 <= 2 {
                        dmg += j2 + 1
                    } else {
                        armor += j2 - 2
                    }

                    if d21Boss(dmg, armor) {
                        if runCost < lowestCost {
                            lowestCost = runCost
                        }
                    } else {
                        if runCost > highestCost {
                            highestCost = runCost
                        }
                    }
                }
            }
        }
    }

    fmt.Println("Part 1: ", lowestCost)
    fmt.Println("Part 2: ", highestCost)
}
