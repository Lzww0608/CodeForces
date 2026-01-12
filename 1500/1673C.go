package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD int = 1_000_000_007
const N int = 40_001

var a []int
var f [N]int

func init() {
	check := func(x int) bool {
		y := 0
		t := x
		for x > 0 {
			y = y*10 + x%10
			x /= 10
		}

		return t == y
	}
	for i := 1; i < N; i++ {
		if check(i) {
			a = append(a, i)
		}
	}

	f[0] = 1
	for _, x := range a {
		for i := x; i < N; i++ {
			f[i] = (f[i] + f[i-x]) % MOD
		}
	}
}

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
	fmt.Fprintln(out, f[n])
}
