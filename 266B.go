package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {

	var n, t int
	in := bufio.NewReader(os.Stdin)
	Fscan(in, &n, &t)
	var s []byte
	Fscan(in, &s)

	f := true
	for i := 0; i < t && f; i++ {
		f = false
		for j := 0; j < len(s)-1; j++ {
			if s[j] == 'B' && s[j+1] == 'G' {
				s[j], s[j+1] = s[j+1], s[j]
				j++
				f = true
			}
		}
	}
	Println(string(s))
	return
}
