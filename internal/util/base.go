package util

import (
    "bufio"
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