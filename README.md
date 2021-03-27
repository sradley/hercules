# Hercules
A more versatile brute-force tool. Out-of-the-box support for multiple protocols and
parallelized connects.

```
Hercules v0.0.1 (c) 2021 by Stephen Radley

Hercules (https://github.com/sradley/hercules) starting at 2021-03-28 02:13:20
[+] 5 threads, 14 login attempts, ~3 attempts per thread
[*] username: root       password: 123456

[+] attack completed, 1 valid credentials(s) found
Hercules (https://github.com/sradley/hercules) finished at 2021-03-28 02:13:20
```

## Table of Contents
 * [Installation](#installation)
   - [Using Go Get](#using-go-get)
   - [Building from Source](#building-from-source)

## Installation

### Using Go Get
If you have a golang environment set up, you can use `go get` to fetch and install the binary in
your `$GOPATH/bin` directory.
```
$ go get github.com/sradley/hercules
```

### Building from Source
The following commands will install dependencies and compile a working binary.
```
$ git clone https://github.com/sradley/hercules.git
$ cd overflow
$ go get && go build
```

You can also use `go install` to install the binary in your `$GOPATH/bin` directory.
```
$ git clone https://github.com/sradley/hercules.git
$ cd overflow
$ go install
```