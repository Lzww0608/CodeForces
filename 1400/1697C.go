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
	var s, t string
	fmt.Fscan(in, &s, &t)
	if strings.Count(s, "b") != strings.Count(t, "b") {
		fmt.Fprintln(out, "NO")
		return
	}
	j := 0
	for i := range s {
		if s[i] == 'b' {
			continue
		}
		for t[j] == 'b' {
			j++
		}

		if s[i] != t[j] || (s[i] == 'a' && i > j) || (s[i] == 'c' && i < j) {
			fmt.Fprintln(out, "NO")
			return
		}
		j++
	}

	fmt.Fprintln(out, "YES")
}
