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

	var n, l, r int
	fmt.Fscan(in, &n, &l, &r)
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &b[i])
	}
	l--
	r--
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		if (i < l || i > r) && a[i] != b[i] {
			fmt.Fprintln(out, "LIE")
			return
		} else {
			cnt[a[i]]++
			cnt[b[i]]--
		}
	}

	for _, x := range cnt {
		if x != 0 {
			fmt.Fprintln(out, "LIE")
			return
		}
	}
	fmt.Fprintln(out, "TRUTH")
}
