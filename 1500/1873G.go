package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	a := strings.Count(s, "A")
	b := strings.Count(s, "B")
	if a == 0 || b == 0 {
		fmt.Fprintln(out, 0)
		return
	}
	if s[0] == 'B' || s[n-1] == 'B' {
		fmt.Fprintln(out, a)
		return
	}

	c := make([]int, 0, n)
	b = 0
	cnt := 0
	for i := 0; i < n; i++ {
		if s[i] == 'A' {
			cnt++
		} else {
			if cnt > 0 {
				c = append(c, cnt)
				cnt = 0
			}
			if s[i+1] == 'A' || s[i-1] == 'A' {
				b++
			}
		}
	}
	c = append(c, cnt)

	if b >= len(c) {
		fmt.Fprintln(out, a)
		return
	}
	fmt.Fprintln(out, a-slices.Min(c))
}
