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
	g := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		g = gcd(a[i], g)
	}
	ans := 0
	for i := 1; i*i <= g; i++ {
		if g%i == 0 {
			ans++
			if i*i != g {
				ans++
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
