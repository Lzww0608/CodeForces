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

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	mx := a[n-1] + 1
	f := make([]int, mx)
	primes := make([][]int, mx)
	for i := 2; i < mx; i++ {
		if len(primes[i]) != 0 {
			continue
		}
		for j := i; j < mx; j += i {
			primes[j] = append(primes[j], i)
		}
	}

	ans := 1
	for _, x := range a {
		cur := 0
		for _, y := range primes[x] {
			cur = max(cur, f[y])
		}
		cur++
		for _, y := range primes[x] {
			f[y] = cur
		}
		ans = max(ans, cur)
	}
	fmt.Fprintln(out, ans)
}
