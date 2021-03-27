package util

import (
    "bufio"
    "fmt"
    "os"
    "time"
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

    t := time.Now()
    fmt.Printf("Hercules (https://github.com/sradley/hercules) starting at %s\n",
        t.Format("2006-01-02 15:04:05"))
}

func PrintFooter() {
    t := time.Now()
    fmt.Printf("Hercules (https://github.com/sradley/hercules) finished at %s\n",
        t.Format("2006-01-02 15:04:05"))
}

func PrintResult(username, password string) {
    fmt.Printf("[*] username: \u001b[1m\u001b[32m%s\u001b[0m\t password: \u001b[1m\u001b[32m%s\u001b[0m\n",
        username, password)
}