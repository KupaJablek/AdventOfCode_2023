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

    var f1 string 
    var f2 string 
    switch day {
    case "01", "02", "03", "04", "05", "06", "08", "10":
        f1 = "p1.py"
        f2 = "p2.py"
    case "09":
        f1 = "p1_p2.py"
    case "07":
        day7_p1()
        f1 = "p2.py"
    case "11":
        D11() 
    case "12":
        d12()
    case "15":
        d15()
    case "16":
        d16()
    case "17":
        d17()
    case "18":
        d18()
    case "14":
        d14()
    case "19":
        d19()
    case "13", "20", "21", "22", "23", "24", "25":
        fmt.Println("Day " + day + " solution is not yet implemented")
    default:
        fmt.Println("Day: " + day + " is not valid day [1-25]")
    }

    if f1 == "" {
        return
    }

    cmd := exec.Command("python3", "./" + day + "/" + f1, "./" + day + "/input.txt")
    cmd2 := exec.Command("python3", "./" + day + "/" + f2, "./" + day + "/input.txt")

    out := &bytes.Buffer{}
    cmd.Stdout = out
    cmd.Run()
    fmt.Println(string(out.Bytes()))

    if f2 != "" {
        fmt.Println("Part 2")
        out2 := &bytes.Buffer{}
        cmd2.Stdout = out2
        cmd2.Run()
        fmt.Println(string(out2.Bytes()))
    }
}
