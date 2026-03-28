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
	var p, q int
	fmt.Fscan(in, &p, &q)
	if p < q || p%q != 0 {
		fmt.Fprintln(out, p)
		return
	}

	x := p
	cnt := make(map[int]int)
	for i := 2; i <= p; i++ {
		if p%i == 0 {
			a := 0
			for p%i == 0 {
				p /= i
				a++
			}
			cnt[i] = a
		}
	}

	ans := 1
	for k, v := range cnt {
		a := 0
		for t := q; t%k == 0; t /= k {
			a++
		}
		ans = max(ans, x/pow(k, v-a+1))
	}
	fmt.Fprintln(out, ans)
}

func pow(a, b int) int {
	res := 1
	for b > 0 {
		if b&1 == 1 {
			res *= a
		}
		a *= a
		b >>= 1
	}
	return res
}
