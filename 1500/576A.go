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
	ans := []int{}
	vis := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		if vis[i] {
			continue
		}
		for j := i * i; j <= n; j += i {
			vis[j] = true
		}
		for j := i; j <= n; j *= i {
			ans = append(ans, j)
		}
	}
	fmt.Fprintln(out, len(ans))
	for _, x := range ans {
		fmt.Fprintf(out, "%d ", x)
	}
}
