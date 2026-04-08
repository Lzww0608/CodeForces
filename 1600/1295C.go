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

const N = 26

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s, t string
	fmt.Fscan(in, &s, &t)

	n := len(s)
	f := make([][N]int, n+1)
	p := [N]int{}
	for i := range p {
		p[i] = n
	}

	for i := n - 1; i >= 0; i-- {
		x := int(s[i] - 'a')
		p[x] = i
		f[i] = p
	}

	ans := 0
	for i := 0; i < len(t); {
		j := 0
		if f[0][int(t[i]-'a')] == n {
			fmt.Fprintln(out, -1)
			return
		}
		for j < n && i < len(t) {
			x := int(t[i] - 'a')
			if f[j][x] == n {
				break
			}
			j = f[j][x] + 1
			i++
		}
		ans++

	}

	fmt.Fprintln(out, ans)
}
