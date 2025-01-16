package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var N int
var Primes []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	N = int(math.Sqrt(1_000_000_000))
	isPrime := make([]int, N)

	for i := 2; i < N; i++ {
		if isPrime[i] == 1 {
			continue
		}
		Primes = append(Primes, i)
		for j := i * i; j < N; j += i {
			isPrime[j] = 1
		}
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}

}

func decompose(x int) (res int) {
	for _, p := range Primes {
		if p*p > x {
			break
		}
		for x%p == 0 {
			x /= p
			res++
		}
	}

	if x > 1 {
		res++
	}
	return
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var a, b, k int
	fmt.Fscan(in, &a, &b, &k)
	m, n := 0, decompose(a)+decompose(b)
	g := gcd(a, b)
	if a == b {
		m = 0
	} else if g == a || g == b {
		m = 1
	} else {
		m = 2
	}

	if k >= m && k <= n {
		if k != 1 || k == 1 && m == 1 {
			fmt.Fprintln(out, "YES")
			return
		}
	}
	fmt.Fprintln(out, "NO")
	return
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
