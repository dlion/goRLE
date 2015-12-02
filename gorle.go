package gorle

import (
	"bytes"
	"fmt"
	"strconv"
)

func Encode(s string) string {
	count := 1
	var o bytes.Buffer
	var i int
	for i = 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			count++
		} else {
			o.WriteString(fmt.Sprintf("%d%c", count, s[i]))
			count = 1
		}
	}
	o.WriteString(fmt.Sprintf("%d%c", count, s[i]))
	return o.String()
}
