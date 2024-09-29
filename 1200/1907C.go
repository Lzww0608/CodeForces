package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		var s string
		fmt.Fscan(in, &s)
		cnt := make([]int, 26)
		for i := range s {
			x := int(s[i]&31) - 1
			cnt[x]++
		}
		mx := slices.Max(cnt)
		if mx*2-n <= 1 {
			fmt.Fprintln(out, n&1)
		} else {
			fmt.Fprintln(out, mx*2-n)
		}

	}
}
