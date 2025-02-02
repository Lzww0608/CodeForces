package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	var n, c int
	fmt.Fscan(in, &n, &c)
	a := make([]int, n)
	p := make([]int, n)
	for i := range a {
		p[i] = i
		fmt.Fscan(in, &a[i])
	}
	sort.Slice(p, func(i, j int) bool {
		return a[p[i]]-c*(p[i]+1) > a[p[j]]-c*(p[j]+1)
	})

	cur := a[0]
	for _, i := range p {
		if i == 0 {
			continue
		}
		if c*(i+1) > cur+a[i] {
			fmt.Fprintln(out, "NO")
			return
		}
		cur += a[i]
	}
	fmt.Fprintln(out, "YES")
}
