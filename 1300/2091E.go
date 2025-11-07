package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 10_000_001

var isPrime [N]bool

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i := 2; i < N; i++ {
		if isPrime[i] {
			continue
		}

		for j := i * i; j < N; j += i {
			isPrime[j] = true
		}
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	ans := 0
	for i := 2; i <= n; i++ {
		if !isPrime[i] {
			ans += n / i
		}
	}
	fmt.Fprintln(out, ans)
}
