# Go Detective

An ohai-like platform sniffer for Go.

## Usage

```go
package main

import "fmt"
import "github.com/kyleterry/detective"

func main() {
    detective.Init()
    d := detective.CollectData()
    fmt.Println(d)
}
```
