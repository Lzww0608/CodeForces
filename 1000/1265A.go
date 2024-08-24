package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		var str string
		fmt.Fscan(in, &str)
		if strings.Contains(str, "aa") || strings.Contains(str, "bb") || strings.Contains(str, "cc") {
			fmt.Fprintln(out, -1)
			continue
		}
		s := []byte(str)
		n := len(s)
		for i := range s {
			if s[i] == '?' {
				a, b := true, true
				if i > 0 {
					if s[i-1] == 'a' {
						a = false
					} else if s[i-1] == 'b' {
						b = false
					}
				}
				if i < n-1 {
					if s[i+1] == 'a' {
						a = false
					} else if s[i+1] == 'b' {
						b = false
					}
				}
				if a {
					s[i] = 'a'
				} else if b {
					s[i] = 'b'
				} else {
					s[i] = 'c'
				}

			}
		}
		fmt.Fprintln(out, string(s))
	}

}
