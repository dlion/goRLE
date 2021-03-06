# goRLE
Didactics version of [Run-Length Encoding](https://en.wikipedia.org/wiki/Run-length_encoding) on strings in Golang

##Install

```
go install github.com/dlion/goRLE
```

##Usage

```go
package main

import (
	"fmt"
	"github.com/dlion/goRLE"
)

func main() {
	value := "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW""
	enc := gorle.Encode(value)
	fmt.Printf("%s\n%s\n", value, enc)
}
```

```
WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW
12W1B12W3B24W1B14W
```

## Author
* Domenico Luciani
* http://dlion.it

## License
MIT
