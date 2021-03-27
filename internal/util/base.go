package util

import (
    "bufio"
    "fmt"
    "os"
)

func ReadLines(file string) ([]string, error) {
    fp, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer fp.Close()

    var lines []string
    sc := bufio.NewScanner(fp)

    for sc.Scan() {
        lines = append(lines, sc.Text())
    }

    return lines, sc.Err()
}

func PrintHeader() {
    fmt.Println("Hercules v0.0.1 (c) 2021 by Stephen Radley\n")
}