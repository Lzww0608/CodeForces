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
	cnt := make([]int, n)
	mx := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		x--
		cnt[x]++
		mx = max(mx, cnt[x])
	}
	y, t := 0, 0
	for _, i := range cnt {
		if i == mx {
			t++
		} else {
			y += i
		}
	}
	fmt.Fprintln(out, y/(mx-1)+t-1)
}
