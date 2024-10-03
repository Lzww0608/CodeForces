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
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		var s string
		fmt.Fscan(in, &s)
		if s[len(s)-1] != 'B' {
			fmt.Fprintln(out, "NO")
			continue next
		}
		cnt := 0
		for _, c := range s {
			if c == 'A' {
				cnt++
			} else {
				cnt--
				if cnt < 0 {
					fmt.Fprintln(out, "NO")
					continue next
				}
			}

		}
		fmt.Fprintln(out, "YES")
	}
}
