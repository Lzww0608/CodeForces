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
	ans := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			ans = min(ans, i)
			if gcd(i, n/i) == 1 {
				ans = 1
				break
			}
		}
	}
	fmt.Fprintln(out, ans)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
