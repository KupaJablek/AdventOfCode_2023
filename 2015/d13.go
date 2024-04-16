package main

import "fmt"

func d13PermuteList(g []string) {
    seen := [][]string{}

    n, i := len(g), 1

    t := make([]string, n)
    for k := range t {
        t[k] = g[k]
    }
    seen = append(seen, t)
    
    c := make([]int, n)
    for i < n {
        if c[i] < i {
            if i % 2 == 0 {
                g[0], g[i] = g[i], g[0]
            } else {
                g[c[i]], g[i] = g[i], g[c[i]]
            }
            temp := make([]string, n)
            for k := range temp {
                temp[k] = g[k]
            }
            seen = append(seen, temp)
            c[i]++
            i = 1
        } else {
            c[i] = 0
            i++ 
        }
    }
    fmt.Println("Possible permutations: ", len(seen))
}

func d13() {
    peeps := []string{"Alice", "Bob", "Carol", "David", "Frank", "George", "Mallory"}
    d13PermuteList(peeps) 
}
