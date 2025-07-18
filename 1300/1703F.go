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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, 0, n)
	b := make([]int, 0, n)

	var x int
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &x)
		if i > x {
			a = append(a, i)
			b = append(b, x)
		}
	}

	sort.Ints(a)
	sort.Ints(b)
	ans := 0

	i := 0
	for _, x := range a {
		for i < len(b) && x >= b[i] {
			i++
		}
		ans += len(b) - i
	}

	fmt.Fprintln(out, ans)
}
