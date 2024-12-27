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

	var n, k, s int
	fmt.Fscan(in, &n, &k, &s)
	cur := 1
	if k > s || k*(n-1) < s {
		fmt.Fprintln(out, "NO")
		return
	}
	fmt.Fprintln(out, "YES")
	for k > 0 {
		t := min(n-1, s-(k-1))
		if cur-t > 0 {
			cur -= t
		} else {
			cur += t
		}
		fmt.Fprint(out, cur, " ")
		k--
		s -= t
	}
}
