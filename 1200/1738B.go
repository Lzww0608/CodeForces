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
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, k)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	if k == 1 {
		fmt.Fprintln(out, "YES")
		return
	}

	for i := 1; i < k-1; i++ {
		if a[i]-a[i-1] > a[i+1]-a[i] {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	t := a[1] - a[0]
	if t*(n-k+1) < a[0] {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
}
