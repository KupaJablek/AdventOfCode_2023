package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
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

type star struct {
    y int
    x int
}


func parseData(d []string) {
    var stars []star
    var eRows []int
    var eCols []int

    for i := 0; i < len(d); i++ {
        rowEmpty := true
        for j := 0; j < len(d[i]); j++ {
            if d[i][j] == '#' {
                rowEmpty = false
                stars = append(stars, star{i, j})

                if !slices.Contains(eCols, j) {
                    eCols = append(eCols, j)
                }
            }
        }
        if rowEmpty {
            eRows = append(eRows, i)
        }
    }

    for i := 0; i < len(eCols); i++ {
        fmt.Printf("%d, ", eCols[i])
    }
    fmt.Println()
}

func D11() {
    //data, err := ParseFile("./11/input.txt")
    data, err := ParseFile("./11/test.txt")
    if err != nil {
        fmt.Println(err.Error())   
        return
    }

    for i := 0; i < len(data); i++ {
        fmt.Println(data[i])
    }
    parseData(data)
}
