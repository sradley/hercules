package cmd

import (
    "os"
    "github.com/spf13/cobra"
    "github.com/sradley/hercules/cmd/http"
)

var root = &cobra.Command{
    Use:          "hercules",
    SilenceUsage: true,
}

func Execute() {
    if err := root.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
    root.AddCommand(http.SubCommand)
    http.Init()
}
