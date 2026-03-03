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

	var n, a, b, k int
	fmt.Fscan(in, &n, &a, &b, &k)
	h := make([]int, n)
	for i := range h {
		fmt.Fscan(in, &h[i])
		h[i] %= a + b
		if h[i] == 0 {
			h[i] = a + b
		}
		h[i] = (h[i]+a-1)/a - 1
	}
	sort.Ints(h)
	ans := 0
	for _, x := range h {
		if x <= k {
			k -= x
			ans++
		} else {
			break
		}
	}

	fmt.Fprintln(out, ans)
}
