package http

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/sradley/hercules/internal/http"
)

var (
    request   string
    users     string
    passwords string
    bad       string
    threads   int
)

var SubCommand = &cobra.Command{
    Use:   "http",
    Short: "Attacks an HTTP-based login, from a given HTTP request",
    Run:   start,
}

func Init() {
    SubCommand.Flags().StringVarP(&request, "request", "r", "",
        "the target HTTP request")
    SubCommand.MarkFlagRequired("request")

    SubCommand.Flags().StringVarP(&users, "users", "u", "",
        "(optional) the list of users to brute force")

    SubCommand.Flags().StringVarP(&passwords, "passwords", "p", "",
        "(optional) the list of passwords to brute force")

    SubCommand.Flags().StringVarP(&bad, "bad", "b", "",
        "(optional) a string only found in a negative response")

    SubCommand.Flags().IntVarP(&threads, "threads", "t", 20,
        "(optional) the number of concurrent threads to utilise")
}

func start(cmd *cobra.Command, argv []string) {
    args, err := http.ParseArgs(request, users, passwords, bad, threads)
    if err != nil {
        fmt.Println(err)
        return
    }

    attempts := len(args.Users) * len(args.Passwords)
    fmt.Printf("[+] %d threads, %d login attempts, ~%d attempts per thread\n",
        args.Threads, attempts, attempts / args.Threads + 1)

    http.Attack(args)

    fmt.Printf("\n[+] attack completed, %d valid credentials(s) found\n", 1)
}
