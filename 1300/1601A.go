package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)

	cnt := [31]int{}
	for i := range a {
		fmt.Fscan(in, &a[i])
		for j := 0; j < 31; j++ {
			if a[i]&(1<<uint(j)) != 0 {
				cnt[j]++
			}
		}
	}

	g := 0
	for _, x := range cnt {
		g = gcd(x, g)
	}

	for i := 1; i <= n; i++ {
		if g%i == 0 {
			fmt.Fprintf(out, "%d ", i)
		}
	}

	fmt.Fprintln(out)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
