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
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	for _, t := range []string{"aa", "aba", "aca", "abca", "acba", "accabba", "abbacca"} {
		if strings.Contains(s, t) {
			fmt.Fprintln(out, len(t))
			return
		}
	}

	fmt.Fprintln(out, -1)
}
