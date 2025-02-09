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
	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
	}

	dg := make([]int, n*2)
	udg := make([]int, n*2)
	f := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			f[j] ^= dg[i+j] ^ udg[i-j+n]
			if f[j]^int(s[i][j]-'0') != 0 {
				ans++
				f[j] ^= 1
				dg[i+j] ^= 1
				udg[i-j+n] ^= 1
			}
		}
	}
	fmt.Fprintln(out, ans)
}
