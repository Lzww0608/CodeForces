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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

	cntA := make([]int, n+1)
	cntB := make([]int, n+1)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		if y == 0 {
			cntA[x]++
		} else {
			cntB[x]++
		}
	}

	ans := 0
	for i := range n + 1 {
		if cntA[i] > 0 && cntB[i] > 0 {
			ans += n - 2
		}
	}

	for i := 1; i < n; i++ {
		if cntA[i] > 0 && cntB[i-1] > 0 && cntB[i+1] > 0 {
			ans++
		}
		if cntB[i] > 0 && cntA[i-1] > 0 && cntA[i+1] > 0 {
			ans++
		}
	}
	fmt.Fprintln(out, ans)
}
