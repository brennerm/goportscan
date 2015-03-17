[![Build Status](https://travis-ci.org/brennerm/goportscan.svg?branch=master)](https://travis-ci.org/brennerm/goportscan)
[![codecov.io](https://codecov.io/github/brennerm/goportscan/coverage.svg?branch=master)](https://codecov.io/github/brennerm/goportscan?branch=master)

##Installation

```bash
$ go get github.com/brennerm/goportscan
```

##Testing
Run tests with
```bash
$ go test
```

##Documentation

###NewPortScanner()
```go
func NewPortScanner(host string) PortScanner
```
returns a PortScanner instance for given __host__

###ScanPortRange()
```go
func (ps PortScanner)ScanPortRange(start_port int, end_port int) map[int]string
```
scans host for ports between __start\_port__ and __end\_port__
resolves service if port is known
returns map with open port as key and service description as value

###ScanPorts()
```go
func (ps PortScanner)ScanPorts(ports []int) map[int]string
```
scans host for ports in __ports__
resolves service if port is known
returns map with open port as key and service description as value

###ScanKnownPorts()
```go
func (ps PortScanner)ScanKnownPorts() map[int]string
```
scans host for known ports
returns map with open port as key and service description as value

##Example Usage

```go
package main

import (
    "fmt"
    "github.com/brennerm/goportscan"
)

func main(){
    ps := goportscan.NewPortScanner("127.0.0.1")

    for port, service := range ps.ScanKnownPorts(){
	fmt.Printf("open port found:  %d %s\n", port, service)
    }

    for port, service := range ps.ScanPorts([]int{80, 123}){
        fmt.Printf("open port found:  %d %s\n", port, service)
    }
   
    // use with caution
    for port, service := range ps.ScanPortRange(1, 200){
        fmt.Printf("open port found:  %d %s\n", port, service)
    }  
}
```

## License

This project is licensed under the MIT open source license.
