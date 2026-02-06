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
	var pos int
	fmt.Fscan(in, &pos)
	pos--

	cnt, p := 0, 0
	n := len(s)
	for i := range n {
		t := n - i
		if pos < t {
			cnt, p = i, pos
			break
		}
		pos -= t
	}

	st := []byte{}
	for i := range s {
		for cnt > 0 && len(st) > 0 && st[len(st)-1] > s[i] {
			st = st[:len(st)-1]
			cnt--
		}
		st = append(st, s[i])
	}

	fmt.Fprintf(out, "%c", st[p])
}
