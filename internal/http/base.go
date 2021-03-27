package http

import (
    "fmt"
    "bufio"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "github.com/sradley/hercules/internal/pool"
)

type job struct {
    request  string
    username string
    password string
    bad      string
}

func Attack(args *Args) {
    attempts := len(args.Users) * len(args.Passwords)
    fmt.Printf("%d threads, %d login attempts, ~%d attempts per thread\n", args.Threads, attempts, attempts / args.Threads + 1)

    for _, user := range args.Users {
        jobs := []pool.Job{}

        for _, pass := range args.Passwords {
            jobs = append(jobs, job{ args.Request, user, pass, args.Bad })
        }

        pool.Start(jobs, args.Threads)
    }
}

func (j job) Run() (*pool.Result, error) {
    j.request = strings.Replace(j.request, "^USER^", j.username, -1)
    j.request = strings.Replace(j.request, "^PASS^", j.password, -1)

    request, err := buildRequest(j.request)
    if err != nil {
        return &pool.Result{ j.username, j.password, false }, err
    }

    client := &http.Client{}

    resp, err := client.Do(request)
    if err != nil {
        return &pool.Result{ j.username, j.password, false }, err
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return &pool.Result{ j.username, j.password, false }, err
    }

    if !strings.Contains(string(body), j.bad) {
        return &pool.Result{ j.username, j.password, false }, nil
    }

    fmt.Printf("\nFOUND: %s : %s\n", j.username, j.password)

    return &pool.Result{ j.username, j.password, true }, nil
}

func buildRequest(request string) (*http.Request, error) {
    r := bufio.NewReader(strings.NewReader(request))
    _request, err := http.ReadRequest(r)
    if err != nil {
        return nil, err
    }

    // TODO: support https
    _request.URL, err = url.Parse("http://" + _request.Host + _request.RequestURI)
    if err != nil {
        return nil, err
    }

    _request.RequestURI = "";

    return _request, nil
}