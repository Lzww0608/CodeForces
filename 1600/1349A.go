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
	suf := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	for i := n - 1; i >= 0; i-- {
		suf[i] = gcd(a[i], suf[i+1])
	}

	ans := 0
	for i, x := range a {
		cur := lcm(x, suf[i+1])
		ans = gcd(cur, ans)
	}

	fmt.Fprintln(out, ans)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
