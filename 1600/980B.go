package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	fmt.Fprintln(out, "YES")
	fmt.Fprintln(out, strings.Repeat(".", n))
	if k&1 == 0 {
		fmt.Fprint(out, ".")
		fmt.Fprint(out, strings.Repeat("#", k/2))
		fmt.Fprintln(out, strings.Repeat(".", n-k/2-1))
		fmt.Fprint(out, ".")
		fmt.Fprint(out, strings.Repeat("#", k/2))
		fmt.Fprintln(out, strings.Repeat(".", n-k/2-1))
	} else {
		k--
		fmt.Fprint(out, ".")
		if k <= n-3 {
			t := (n - 3 - k) / 2
			fmt.Fprint(out, strings.Repeat(".", t))
			fmt.Fprint(out, strings.Repeat("#", k+1))
			fmt.Fprintln(out, strings.Repeat(".", t+1))
			fmt.Fprintln(out, strings.Repeat(".", n))
		} else {
			fmt.Fprint(out, strings.Repeat("#", n-2))
			fmt.Fprintln(out, ".")
			k -= n - 3
			t := (n - k) / 2
			fmt.Fprint(out, strings.Repeat(".", t))
			fmt.Fprint(out, strings.Repeat("#", k/2))
			fmt.Fprint(out, ".")
			fmt.Fprint(out, strings.Repeat("#", k/2))
			fmt.Fprintln(out, strings.Repeat(".", t))
		}
	}
	fmt.Fprintln(out, strings.Repeat(".", n))
}
