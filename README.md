# Detective

An ohai-like platform sniffer for Go.

## Usage

```go
package main

import "fmt"
import "github.com/kyleterry/detective"

func main() {
    detective.Init()
    d := detective.CollectAllMetrics()
    fmt.Println(d)
}
```

## Testing

I'm not sure how to test this with so many platforms available. If anyone has
any ideas, please shoot me an email.
