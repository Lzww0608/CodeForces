package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 998244353

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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	f := make([]int, 4)
	for _, x := range a {
		if x == 1 {
			f[1] = (f[1] + 1) % MOD
		} else if x == 2 {
			f[2] = (f[2]*2 + f[1]) % MOD
		} else {
			f[3] = (f[3] + f[2]) % MOD
		}
	}

	fmt.Fprintln(out, f[3])
}
