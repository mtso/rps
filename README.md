# rps [![GoDoc](https://godoc.org/github.com/mtso/rps?status.svg)](https://godoc.org/github.com/mtso/rps)
Rock paper scissors chooser.

## Install
```
$ go get github.com/mtso/rps
```

## Use
```go
package main

import (
    "fmt"

    "github.com/mtso/rps"
)

func main() {
    choice := rps.Roll()
    fmt.Println(choice) // [rock|paper|scissors]
    
    choice = rps.Roll([]string{
        "foo",
        "bar",
        "baz",
    })
    fmt.Println(choice) // [foo|bar|baz]
    
    rps.Set([]string{
        "foo",
        "bar",
        "baz",
    })
    choice = rps.Roll()
    fmt.Println(choice) // [foo|bar|baz]
}
```
