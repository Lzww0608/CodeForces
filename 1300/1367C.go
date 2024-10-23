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
	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)
	pos := []int{}
	for i := range s {
		if s[i] == '1' {
			pos = append(pos, i)
		}
	}
	if len(pos) == 0 {
		fmt.Fprintln(out, (n+k)/(k+1))
		return
	}
	pos = append(pos, n+k)
	ans := pos[0] / (k + 1)
	l := pos[0]
	for _, x := range pos[1:] {
		d := x - l - k - 1
		l = x
		ans += d / (k + 1)
	}
	fmt.Fprintln(out, ans)
}
