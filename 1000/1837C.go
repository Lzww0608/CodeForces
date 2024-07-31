package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		var s string
		fmt.Fscan(in, &s)
		n := len(s)
		str := []byte(s)
		for i, c := range s {
			if c == '?' {
				if i > 0 && str[i-1] == '1' {
					str[i] = '1'
				} else if i < n-1 && str[i+1] == '1' {
					str[i] = '1'
				} else {
					str[i] = '0'
				}

			}
		}
		fmt.Fprintln(out, string(str))
	}
}
