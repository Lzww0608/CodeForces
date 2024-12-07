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
	var n, x int
	fmt.Fscan(in, &n)
	fa := make([]int, n+1)
	depth := make([]int, n+1)
	root := -1
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &fa[i])
		if fa[i] == i {
			root = i
		}
	}

	fmt.Fscan(in, &x)
	ok := x == root
	depth[x] = 1
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &x)
		if ok {
			if depth[fa[x]] == 0 {
				ok = false
			}
			depth[x] = i
		}
	}

	if !ok {
		fmt.Fprintln(out, -1)
		return
	}

	for i := 1; i <= n; i++ {
		v := depth[i]
		fmt.Fprint(out, v-depth[fa[i]], " ")
	}
	fmt.Fprintln(out)
}
