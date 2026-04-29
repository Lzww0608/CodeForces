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
	var n, c int
	fmt.Fscan(in, &n, &c)
	a := make([]int, n)
	mx, p := 0, -1
	for i := range a {
		fmt.Fscan(in, &a[i])
		if i > 0 && a[i] > mx {
			mx = a[i]
			p = i
		}
	}

	if n == 1 {
		fmt.Fprintln(out, 0)
		return
	}

	all := max(a[0]+c, mx)
	if all == a[0]+c {
		p = 0
	}
	sum := c
	for i, x := range a {
		ans := 0
		if i == p {
			ans = 0
		} else if sum+x >= mx {
			ans = i
		} else {
			ans = i + 1
		}

		fmt.Fprintf(out, "%d ", ans)
		sum += x
	}

	fmt.Fprintln(out)
}
