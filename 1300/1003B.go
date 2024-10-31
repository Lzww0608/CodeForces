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

	var a, b, x int
	fmt.Fscan(in, &a, &b, &x)
	if a > b {
		solve(out, a, b, x, '0', '1')
	} else {
		solve(out, b, a, x, '1', '0')
	}
}

func solve(out *bufio.Writer, a, b, x int, c, d byte) {
	n := a + b
	s := make([]byte, n)

	id := 0
	for x > 1 {
		if id&1 == 0 {
			s[id] = c
			a--
		} else {
			s[id] = d
			b--
		}
		id++
		x--
	}
	if id > 0 && s[id-1] == c {
		for b > 0 {
			b--
			s[id] = d
			id++
		}
	} else {
		for a > 0 {
			a--
			s[id] = c
			id++
		}
	}
	for a > 0 {
		a--
		s[id] = c
		id++
	}
	for b > 0 {
		b--
		s[id] = d
		id++
	}

	fmt.Fprintln(out, string(s))
}
