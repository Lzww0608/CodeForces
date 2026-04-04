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
	var p, q int64
	fmt.Fscan(in, &p, &q)
	if p%q != 0 {
		fmt.Fprintln(out, p)
		return
	}

	// Factorize q (q <= 1e9), then try removing each prime factor from p
	// until the number is no longer divisible by q.
	primes := make([]int64, 0)
	tmp := q
	for i := int64(2); i*i <= tmp; i++ {
		if tmp%i == 0 {
			primes = append(primes, i)
			for tmp%i == 0 {
				tmp /= i
			}
		}
	}
	if tmp > 1 {
		primes = append(primes, tmp)
	}

	ans := int64(1)
	for _, prime := range primes {
		cur := p
		for cur%q == 0 {
			cur /= prime
		}
		if cur > ans {
			ans = cur
		}
	}
	fmt.Fprintln(out, ans)
}
