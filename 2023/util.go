package main

import (
	"bufio"
    "os"
	"errors"
)

func ParseFile(p string) ([]string, error) {
    file, err := os.Open(p) 
    if err != nil {
        return nil, errors.New("error: " + err.Error())        
    }

    var data []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
       data = append(data, scanner.Text()) 
    }

    return data, nil
}

func abs(i int) int {
    if i >= 0 {
        return i
    }
    return i * -1
}

func min(lhs, rhs int) int {
    if lhs < rhs {
        return lhs
    }
    return rhs
}

func max(lhs, rhs int) int {
    if lhs < rhs {
        return rhs
    }
    return lhs
}
