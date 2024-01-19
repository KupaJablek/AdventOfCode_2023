package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
    args := os.Args
    if len(args) == 1 {
        return
    }

    f1 := ""
    f2 := ""
    p2 := true

    day := args[1]
    if len(day) == 1 {
        day = "0" + day
    }
    fmt.Println("Running Day", day)

    switch day {
        case "01":
            f1 = "aoc_d1_q1.py"
            f2 = "aoc_d1_q2.py"
        case "02":
            f1 = "aoc_d2_p1.py"
            f2 = "aoc_d2_p2.py"
        case "03":
            f1 = "aoc_d3_p1.py"
            f2 = "aoc_d3_p2.py"
        case "04":
            f1 = "aoc_d4_p1.py"
            f2 = "aoc_d4_p2.py"
        case "05":
            f1 = "aoc_d5_p1.py"
            f2 = "d5_p2.py"
        case "06", "08", "10":
            f1 = "p1.py"
            f2 = "p2.py"
        case "09":
            f1 = "p1_p2.py"
            p2 = false
        case "11":
            D11() 
            p2 = false
        case "07":
            // fix later
            //day7_p1() 
            f1 = "p2.py"
        default:
            fmt.Println("Day: " + day + " is no valid/implemented")
            return
        }

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
