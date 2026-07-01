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
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s string
	fmt.Fscan(in, &s)
	n := len(s)

	for d := n - n&1; d > 0; d-- {
		cnt := 0
		for i := range n - d {
			if s[i] == s[i+d] || s[i] == '?' || s[i+d] == '?' {
				if cnt++; cnt >= d {
					fmt.Fprintln(out, d*2)
					return
				}
			} else {
				cnt = 0
			}
		}
	}

	fmt.Fprintln(out, 0)
}
