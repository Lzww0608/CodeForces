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
	p := make([]int, n)
	a := make([]int, n)
	for i := range p {
		fmt.Fscan(in, &p[i])
		p[i]--
		a[p[i]] = i
	}

	m := make([]bool, n)
	for i := 0; i < n; i++ {
		if m[i] {
			continue
		}

		j := a[i]
		for j < n {
			m[j] = true
			if j+1 == n || m[j+1] {
				break
			}
			if p[j+1] != p[j]+1 {
				fmt.Fprintln(out, "NO")
				return
			}
			j++
		}
	}

	fmt.Fprintln(out, "YES")
}
