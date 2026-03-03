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

	var q int
	for fmt.Fscan(in, &q); q > 0; q-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s, t string
	fmt.Fscan(in, &s, &t)
	if len(t) > len(s) {
		fmt.Fprintln(out, "NO")
		return
	}

	i, j := len(s)-1, len(t)-1
	for ; i >= 0 && j >= 0; j-- {
		for i >= 0 {
			if s[i] == t[j] {
				if i > 0 {
					i--
				}
				break
			}
			i -= 2
		}
	}

	if i < 0 {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
	}
}
