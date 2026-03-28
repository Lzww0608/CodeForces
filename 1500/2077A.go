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
	b := make([]int, n*2)
	for i := range b {
		fmt.Fscan(in, &b[i])
	}
	sort.Ints(b)
	a := make([]int, n*2+1)
	sum := 0

	for i := 0; i < n-1; i++ {
		a[i*2+1] = b[i]
		sum -= b[i]
	}
	for i := 0; i <= n; i++ {
		a[i*2] = b[n+i-1]
		sum += b[n+i-1]
	}
	a[n*2-1] = sum
	for _, x := range a {
		fmt.Fprintf(out, "%d ", x)
	}
	fmt.Fprintln(out)

}
