package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://codeforces.com/problemset/problem/1513/C

const MOD int = 1_000_000_007
const N int = 200_001

var f [N][10]int

func init() {
	for i := range 10 {
		f[0][i] = 1
	}

	for i := 1; i < N; i++ {
		for j := 0; j < 9; j++ {
			f[i][j] = f[i-1][j+1]
		}
		f[i][9] = (f[i-1][0] + f[i-1][1]) % MOD
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
	var n, m int
	fmt.Fscan(in, &n, &m)
	ans := 0
	for n > 0 {
		ans = (ans + f[m][n%10]) % MOD
		n /= 10
	}
	fmt.Fprintln(out, ans)
}
