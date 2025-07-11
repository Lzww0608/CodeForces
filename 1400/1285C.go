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
	ans := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 && gcd(i, n/i) == 1 {
			ans = i
		}
	}
	fmt.Fprintln(out, ans, n/ans)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
