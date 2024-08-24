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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)
	b := []int{}
	c := []int{}

	for i := 0; i < n; i += 2 {
		if len(b) > 0 && a[i] == b[len(b)-1] {
			fmt.Fprintln(out, "NO")
			return
		}
		b = append(b, a[i])
	}
	j := n - 1
	if n&1 == 1 {
		j--
	}
	for ; j >= 0; j -= 2 {
		if len(c) > 0 && a[j] == c[len(c)-1] {
			fmt.Fprintln(out, "NO")
			return
		}
		c = append(c, a[j])
	}
	fmt.Fprintln(out, "YES")
	fmt.Fprintln(out, len(b))
	for _, x := range b {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
	fmt.Fprintln(out, len(c))
	for _, x := range c {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)
}
