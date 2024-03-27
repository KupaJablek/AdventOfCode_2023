package main

import (
	"fmt"
	"slices"
)

var buckets = []int{ 50, 44, 11, 49, 42, 46, 18, 32, 26, 40, 21, 7, 18, 43, 10, 47, 36, 24, 22, 40 }

func d17() {
    //amount := 150
    slices.Sort(buckets)

    fmt.Println(buckets)
    permute(buckets)
}

func permute(nums []int) [][]int {
    //https://en.wikipedia.org/wiki/Heap%27s_algorithm
    ret := [][]int{}

    n, i := len(nums), 1

    temp := make([]int, n)
    for k := range temp {
        temp[k] = nums[k]
    }
    ret = append(ret, temp)
    c := make([]int, n)

    for i < n {
        if c[i] < i {
            if i % 2 == 0 {
                nums[0], nums[i] = nums[i], nums[0]
            } else {
                nums[c[i]], nums[i] = nums[i], nums[c[i]]
            }
            temp := make([]int, n)
            for k := range temp {
                temp[k] = nums[k]
            }
            ret = append(ret, temp)
            c[i]++
            i = 1
        } else {
            c[i] = 0
            i++ 
        }
    }

    return ret
}
