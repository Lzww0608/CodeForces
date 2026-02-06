package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)

	type pair struct {
		x, y int
	}
	m := make(map[pair]int)
	m[pair{0, 0}] = -1

	x, y := 0, 0
	ans := n + 1
	l, r := -1, -1
	for i, c := range s {
		if c == 'L' {
			x--
		} else if c == 'R' {
			x++
		} else if c == 'U' {
			y++
		} else {
			y--
		}

		if v, ok := m[pair{x, y}]; ok {
			if i-v < ans {
				ans = i - v
				l, r = v+2, i+1
			}
		}
		m[pair{x, y}] = i
	}

	if ans > n {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, l, r)
	}
}
