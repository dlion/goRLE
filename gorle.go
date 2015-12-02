package main

import (
	"bytes"
	"fmt"
)

func main() {
	stringa := "WWWWWWWWWWWWBWWWWWWWWWWWWBBBWWWWWWWWWWWWWWWWWWWWWWWWBWWWWWWWWWWWWWW"
	enc := encode(stringa)
	dec := decode(enc)
	fmt.Printf("ENC: %s\nDEC: %s\n", enc, dec)
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
	i := 0
	for i < len(s)-1 {
		if int(s[i]) >= 0 
		for j := 0; j < int(s[i]); j++ {
			o.WriteString(fmt.Sprintf("%c", s[i+1]))
		}
	}
	return o.String()
}
