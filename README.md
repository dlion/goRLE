# goRLE
Didactics version of Run-Length Encoding on strings in Golang

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
	dec := gorle.Decode(enc)
	fmt.Printf("%s\n%s\n%s\n", value, enc, dec)
}
```

```
WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW
12W1B12W3B24W1B14W
WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW
```

## Author
* Domenico Luciani
* http://dlion.it

## License
MIT
