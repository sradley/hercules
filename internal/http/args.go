package http

import (
    "io/ioutil"
    "github.com/sradley/hercules/internal/util"
)

type Args struct {
    Request   string
    Users     []string
    Passwords []string
    Bad       string
    Threads   int
}

func ParseArgs(request, users, passwords, bad string, threads int) (*Args, error) {
    _request, err := ioutil.ReadFile(request)
    if err != nil {
        return nil, err
    }

    _users, err := util.ReadLines(users)
    if err != nil && users != "" {
        return nil, err
    }

    _passwords, err := util.ReadLines(passwords)
    if err != nil && passwords != "" {
        return nil, err
    }

    return &Args{ string(_request), _users, _passwords, bad, threads }, nil
}
