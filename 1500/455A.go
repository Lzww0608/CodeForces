package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 100_001

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x int
	fmt.Fscan(in, &n)
	cnt := make([]int, N)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		cnt[x]++
	}

	f := make([]int, N)
	for i := 1; i < N; i++ {
		f[i] = max(f[max(0, i-2)]+i*cnt[i], f[i-1])
	}

	fmt.Fprintln(out, f[N-1])
}
