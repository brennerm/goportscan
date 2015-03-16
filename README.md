Installation
--------------

```bash
$ go get github.com/brennerm/goportscan
```

Example Usage
----------------

```go
package main

import (
    "fmt"
    "github.com/brennerm/goportscan"
)

func main(){

    ps := portscan.NewPortScanner("127.0.0.1")

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
