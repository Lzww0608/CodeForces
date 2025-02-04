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
	b := make([]int, n)
	sum := 0
	for i := range b {
		fmt.Fscan(in, &b[i])
		sum += b[i]
	}
	if sum%((n+1)*n/2) != 0 {
		fmt.Fprintln(out, "NO")
		return
	}
	sum /= (n + 1) * n / 2
	a := make([]int, n)
	for i := 0; i < n; i++ {
		x := b[(i-1+n)%n] - b[i] + sum
		if x <= 0 || x%n != 0 {
			fmt.Fprintln(out, "NO")
			return
		}
		a[i] = x / n
	}
	fmt.Fprintln(out, "YES")
	for _, x := range a {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
