package main

import (
	"bytes"
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
    
    p2 := true
    f1 := "aoc_d" + string(day[1])
    f2 := "aoc_d" + string(day[1])
    switch day {
        case "01":
            f1 += "_q1.py"
            f2 += "_q2.py"
        case "02", "03", "04":
            f1 += "_p1.py"
            f2 += "_p2.py"
        case "05":
            f1 += "_p1.py"
            f2 = "d5_p2.py"
        case "06", "08", "10":
            f1 = "p1.py"
            f2 = "p2.py"
        case "09":
            f1 = "p1_p2.py"
            p2 = false
        case "07":
            day7_p1()
            f1 = "p2.py"
            p2 = false
        case "11":
            D11() 
            return
        case "12":
            d12()
            return
        case "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25":
            fmt.Println("Day " + day + " solution is not yet implemented")
            return
        default:
            fmt.Println("Day: " + day + " is not valid day [1-25]")
            return
        }

    fmt.Println("Day #", day)
    cmd := exec.Command("python3", "./" + day + "/" + f1, "./" + day + "/input.txt")
    cmd2 := exec.Command("python3", "./" + day + "/" + f2, "./" + day + "/input.txt")

    out := &bytes.Buffer{}
    cmd.Stdout = out
    cmd.Run()
    fmt.Println(string(out.Bytes()))

    if p2 {
        fmt.Println("Part 2")
        out2 := &bytes.Buffer{}
        cmd2.Stdout = out2
        cmd2.Run()
        fmt.Println(string(out2.Bytes()))
    }
}
