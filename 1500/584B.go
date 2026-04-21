package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 1_000_000_007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := quickPow(3, n*3)
	b := quickPow(7, n)
	fmt.Fprintln(out, (a-b+MOD)%MOD)
}

func quickPow(a, r int) int {
	res := 1
	for r > 0 {
		if r&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		r >>= 1
	}
	return res
}
