package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 100_001
const MOD int = 1_000_000_007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, k int
	fmt.Fscan(in, &t, &k)
	f := make([]int, N)
	f[0] = 1
	for i := 1; i < N; i++ {
		f[i] = f[i-1]
		if i-k >= 0 {
			f[i] = (f[i] + f[i-k]) % MOD
		}
	}
	for i := 1; i < N; i++ {
		f[i] = (f[i] + f[i-1]) % MOD
	}

	var l, r int
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &l, &r)
		fmt.Fprintln(out, (f[r]-f[l-1]+MOD)%MOD)
	}

}
