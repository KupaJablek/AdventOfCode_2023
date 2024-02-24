package main

import (
	"container/heap"
	"fmt"
	"slices"
	"strconv"
)

type PriorityQueue []*Item
func (piq PriorityQueue) Len() int {
    return len(piq)
}

func (piq PriorityQueue) Less(i, j int) bool {
   return piq[i].cost > piq[j].cost
}

func (piq PriorityQueue) Swap(i, j int) {
   piq[i], piq[j] = piq[j], piq[i]
   piq[i].index = i
   piq[j].index = j
}

func (piq *PriorityQueue) Push(x interface{}) {
   n := len(*piq)
   item := x.(*Item)
   item.index = n
   *piq = append(*piq, item)
}

func (piq *PriorityQueue) Pop() interface{} {
   old := *piq
   n := len(old)
   item := old[n-1]
   item.index = -1 
   *piq = old[0 : n-1]
   return item
}

type Point struct {
    r, c, dr, dc, n int
}

type Item struct {
    cost, r, c, dr, dc, n, index int
}

func search(m [][]int) int {
    dirs := [][]int{{0,1}, {1,0}, {0,-1}, {-1, 0}}

    pq := make(PriorityQueue, 0)
    heap.Push(&pq, &Item{0,0,0,0,0,0,0})

    seen := []Point{}
    sum := 0
    for len(pq) > 0 {
        // pop first value
        s := heap.Pop(&pq).(*Item)
        fmt.Println(s)

        if s.r == len(m) - 1 && s.c == len(m[0]) - 1 {
            return s.cost
        }

        p := Point{s.r, s.c, s.dr, s.dc, s.n}
        if slices.Contains(seen, p) {
            continue
        }
        seen = append(seen, p)
        
        // check if less than 3 moves in one direction
        if s.n < 3 && s.dr != 0 && s.dc != 0 {
            fmt.Println(s)
            nr := s.r + s.dr
            nc := s.c + s.dc
            if nr >= 0 && nr < len(m) && nc >= 0 && nc < len(m[0]) {
                heap.Push(&pq, &Item{s.cost + m[nr][nc], nr, nc, s.dr, s.dc, s.n + 1, 0})
            }
        }

        // check all directions
        for i := 0; i < 4; i++ {
            p1 := [2]int{dirs[i][0], dirs[i][1]}
            if p1 != [2]int{s.r, s.c} && p1 != [2]int{-s.r, -s.c} {
                nr := s.r + p1[0]
                nc := s.c  + p1[1]
                if nr >= 0 && nr < len(m) && nc >= 0 && nc < len(m[0]) {
                    heap.Push(&pq, &Item{s.cost + m[nr][nc], nr, nc, p1[0], p1[1], 1, 0})
                }
            }
        }
    }

    return sum
}

func d17() {
    //data, err := ParseFile("./17/input.txt")
    data, err := ParseFile("./17/test.txt")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    m := make([][]int, len(data)) 
    for i := 0; i < len(data); i++ {
        s := make([]int, len(data[0]))
        for k := 0; k < len(data[i]); k++ {
            s[k], _ = strconv.Atoi(string(data[i][k]))
        }
        m[i] = s
    }

    fmt.Println("Part 1: ", search(m))
}





