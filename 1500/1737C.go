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
	r := make([]int, 3)
	c := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Fscan(in, &r[i], &c[i])
	}
	u, v := 0, 0
	fmt.Fscan(in, &u, &v)

	rr, cc := r[0], c[0]
	if r[0] != r[1] {
		rr = r[2]
	}
	if c[0] != c[1] {
		cc = c[2]
	}

	if (rr == 1 || rr == n) && (cc == 1 || cc == n) {
		if rr == u || cc == v {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	} else if (rr+cc)&1 == (u+v)&1 {
		if rr&1 == u&1 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	} else {
		fmt.Fprintln(out, "YES")
	}
}
