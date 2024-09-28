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
		pos := 0
		for i := n - 2; i >= 0; i-- {
			if s[i] == 'a' {
				pos = i
				break
			}
		}
		if pos == 0 {
			fmt.Fprintln(out, s[:1], s[1:n-1], s[n-1:])
		} else {
			fmt.Fprintln(out, s[:pos], s[pos:pos+1], s[pos+1:])
		}

	}

}
