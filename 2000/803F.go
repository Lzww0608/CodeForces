package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 100_005
const MOD int = 1_000_000_007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = 1
	}

	var x int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		a[x] = a[x] * 2 % MOD
	}

	for i := 1; i < N; i++ {
		for j := i * 2; j < N; j += i {
			a[i] = a[i] * a[j] % MOD
		}

	}

	b[1] = 1
	for i := 1; i < N; i++ {
		for j := i * 2; j < N; j += i {
			b[j] = (b[j] - b[i] + MOD) % MOD
		}
	}

	ans := 0
	for i := 1; i < N; i++ {
		ans = (ans + b[i]*(a[i]-1)%MOD) % MOD
	}
	fmt.Fprintln(out, ans)
}
