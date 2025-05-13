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

	var n int
	fmt.Fscan(in, &n)
	ans := n
	for i := 2; i <= n; i++ {
		ans = (ans - 1) * i % MOD
	}
	fmt.Fprintln(out, ans)
}
