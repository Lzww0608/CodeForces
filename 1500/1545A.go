package main

import (
	"bufio"
	"fmt"
	"io"
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

func solve(in io.Reader, out io.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	b := []int{}
	for i := range a {
		fmt.Fscan(in, &a[i])
		if i&1 == 1 {
			b = append(b, a[i])
		}
	}

	sort.Ints(a)
	sort.Ints(b)
	for i := range b {
		if b[i] != a[i*2+1] {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
}
