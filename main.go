package main

import (
    "github.com/sradley/hercules/cmd"
    "github.com/sradley/hercules/internal/util"
)

func main() {
    util.PrintHeader()
    cmd.Execute()
    util.PrintFooter()
}
