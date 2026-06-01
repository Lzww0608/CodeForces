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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	mx := 0
	for i := range n {
		cnt := make(map[[2]int]int)

		for j := i + 1; j < n; j++ {
			dy := a[j] - a[i]
			dx := j - i
			g := gcd(abs(dy), abs(dx))
			dy /= g
			dx /= g
			cnt[[2]int{dy, dx}]++
			if cnt[[2]int{dy, dx}] > mx {
				mx = cnt[[2]int{dy, dx}]
			}
		}
	}

	fmt.Fprintln(out, n-mx-1)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
