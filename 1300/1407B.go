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
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		cnt[x]++
	}
	g := 0
	for len(cnt) > 0 {
		cur, x := 0, -1
		for k := range cnt {
			t := gcd(k, g)
			if t > cur {
				cur = t
				x = k
			}
		}
		for i := 0; i < cnt[x]; i++ {
			fmt.Fprint(out, x, " ")
		}
		delete(cnt, x)
		g = cur
	}
	fmt.Fprintln(out)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
