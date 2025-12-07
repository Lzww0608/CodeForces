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
	ans := []int{1}

	p := 1
	for i := 2; i < n; i++ {
		if gcd(i, n) == 1 {
			ans = append(ans, i)
			p = p * i % n
		}
	}

	d := 0
	if p != 1 {
		d = 1
	}
	fmt.Fprintln(out, len(ans)-d)
	for _, x := range ans {
		if p != 1 && x == p {
			continue
		}
		fmt.Fprintf(out, "%d ", x)
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
