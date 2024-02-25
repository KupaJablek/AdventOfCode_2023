package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getArrangements(spring string, nums []int) int {
    if spring == "" {
        if len(nums) == 0 {
            return 1
        }
        return 0
    }

    if len(nums) == 0 {
        if strings.Contains(spring, "#") {
            return 1
        } 
        return 0
    }

    total := 0
    if strings.Contains(".?", string(spring[0])) {
        total += getArrangements(spring[1:], nums)
    }

    if strings.Contains(".#", string(spring[0])) {
        if nums[0] <= len(spring) && !strings.Contains(spring[:nums[0]], ".") && (len(spring) != nums[0] || spring[nums[0]] != '#') {
            total += getArrangements(spring[nums[0] + 1:], nums[1:])
        }
    }

    return total
}

func d12() {
    fmt.Println("Day 12")
    //data, err := ParseFile("./12/input.txt")
    data, err := ParseFile("./12/test.txt")
    if err != nil {
        fmt.Println(err.Error())   
        return
    }

    total := 0
    for d := range data {
        s := strings.Split(data[d], " ")
        n := strings.Split(s[1], ",")
        nums := make([]int, len(n))

        for i := range n {
            nums[i], _ = strconv.Atoi(n[i])
        }

        total += getArrangements(s[0], nums)
    }

    fmt.Println("Part 1:", total)
    fmt.Println("Part 2:", total)
}
