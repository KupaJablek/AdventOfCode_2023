package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
    args := os.Args
    if len(args) != 2 {
        fmt.Println("Usage: go run . [day]")
        return
    }

    day := args[1]
    if len(day) == 1 {
        day = "0" + day
    }
    switch day {
    case "01":
        d1()
    case "02":
        d2()
    case "03":
        d3()
    case "05":
        d5()
    case "04":
        d4()
    case "06":
        d6()
    case "07", "09":
    case "08":
        d8()
    case "10":
        d10()
    case "11":
        d11()
    case "17":
        //d17()
        cmd := exec.Command("python3", "./d17.py")
        out := &bytes.Buffer{}
        cmd.Stdout = out
        cmd.Run()

        fmt.Print(string(out.Bytes()))

    case "13", "16", "18", "19", "20":
    case "14":
        d14()
    case "15":
        d15()
    case "12":
        d12()
    case "21", "22", "23", "24", "25":
    default:
        fmt.Println("Day: " + day + " is not valid day [1-25]")
    }
}

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
