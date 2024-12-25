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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	vis := make([]bool, n+1)
	ans := 0
	id := []int{}
	for i := 0; i < n; i++ {
		if !vis[a[i]] {
			ans++
			b[i] = a[i]
			vis[a[i]] = true
		} else {
			id = append(id, i)
		}
	}

	c := []int{}
	for x := n; x > 0; x-- {
		if !vis[x] {
			c = append(c, x)
		}
	}

	if len(c) == 1 && b[c[0]-1] == 0 {
		for i := 0; i < n; i++ {
			if b[i] == a[c[0]-1] {
				b[i], b[c[0]-1] = c[0], b[i]
				break
			}
		}

	} else {
		pos := -1
		for i := range c {
			if c[i] == id[i]+1 {
				pos = i
			} else {
				b[id[i]] = c[i]
			}
		}

		if pos != -1 {
			if pos == 0 {
				b[id[len(id)-1]], b[id[pos]] = c[pos], b[id[len(id)-1]]
			} else {
				b[id[0]], b[id[pos]] = c[pos], b[id[0]]
			}
		}

	}

	fmt.Fprintln(out, ans)
	for _, x := range b {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out, " ")
}
