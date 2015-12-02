package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	stringa := "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWWWBBBBBBBBBWWW"
	enc := encode(stringa)
	dec := decode(enc)
	fmt.Printf("STR: %s\nENC: %s\nDEC: %s\n", stringa, enc, dec)
}

func encode(s string) string {
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

func decode(s string) string {
	var o bytes.Buffer
	var count string
	for i := 0; i < len(s); i++ {
		val, _ := strconv.Atoi(string(s[i]))
		if val > 0 && val <= 9 {
			if countN, err := strconv.Atoi(count); err == nil {
				count = fmt.Sprintf("%d%d", countN, val)
			} else {
				count = fmt.Sprintf("%d", val)
			}
		} else {
			count2, _ := strconv.Atoi(count)
			for j := 0; j < count2; j++ {
				o.WriteString(fmt.Sprintf("%c", s[i]))
			}
			count = ""
		}
	}
	return o.String()
}
