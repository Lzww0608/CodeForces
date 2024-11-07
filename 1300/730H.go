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

	var n, m int
	fmt.Fscan(in, &n, &m)
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}
	a := make([]int, m)
	for i := range a {
		fmt.Fscan(in, &a[i])
		a[i]--
	}
	d := len(s[a[0]])
	for _, x := range a {
		if d != len(s[x]) {
			fmt.Fprintln(out, "No")
			return
		}
	}
	t := make([]byte, d)
	for i := 0; i < d; i++ {
		c := s[a[0]][i]
		t[i] = c
		for _, x := range a {
			if c != s[x][i] {
				t[i] = '?'
				break
			}
		}
	}

	id := 0
	for i := range s {
		if len(s[i]) != d {
			continue
		} else if id < m && a[id] == i {
			id++
			continue
		}

		f := true
		for j := range s[i] {
			if t[j] != '?' && s[i][j] != t[j] {
				f = false
				break
			}
		}
		if f {
			fmt.Fprintln(out, "No")
			return
		}
	}
	fmt.Fprintln(out, "Yes")
	fmt.Fprintln(out, string(t))
	return
}
