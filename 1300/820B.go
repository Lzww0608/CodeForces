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

	var n, k int
	fmt.Fscan(in, &n, &k)
	t := float64(180) / float64(n)
	ans := 3
	for i := 3; i <= n; i++ {
		if abs(float64(i-2)*t-float64(k)) < abs(float64(ans-2)*t-float64(k)) {
			ans = i
		}
	}
	fmt.Fprintln(out, 2, 1, ans)
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}

	return x
}
