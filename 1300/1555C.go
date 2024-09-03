package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n, x int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		pre1 := make([]int, n+1)
		pre2 := make([]int, n+1)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			pre1[i+1] = pre1[i] + x
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			pre2[i+1] = pre2[i] + x
		}
		ans := math.MaxInt32
		for i := 0; i < n; i++ {
			ans = min(ans, max(pre2[i], pre1[n]-pre1[i+1]))
		}
		fmt.Fprintln(out, ans)
	}
}
