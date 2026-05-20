package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD int = 998_244_353

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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)
	ans := 1
	for i := 2; i <= n; i++ {
		ans = ans * i % MOD
	}

	if a[n-1] == a[n-2] {
		fmt.Fprintln(out, ans)
		return
	}
	k := 0
	for i := n - 2; i >= 0; i-- {
		if a[i] == a[n-1]-1 {
			k++
		} else {
			break
		}
	}

	t := 1
	for i := 2; i <= n; i++ {
		if i != k+1 {
			t = t * i % MOD
		}
	}

	fmt.Fprintln(out, (ans-t+MOD)%MOD)
}
