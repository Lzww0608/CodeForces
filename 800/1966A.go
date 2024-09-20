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

	var t, n, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		x := 0
		freq := map[int]int{}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			freq[x]++
		}
		mx := 0
		for _, x := range freq {
			mx = max(x, mx)
		}
		ans := min(k-1, n)
		if mx < k {
			ans = n
		}
		fmt.Fprintln(out, ans)
	}
}
