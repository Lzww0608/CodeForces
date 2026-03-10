package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 2_000_001
const M int = 1500

var ans [N]int
var cur int = 1

func init() {
	f := make([][]int, M)
	for i := range f {
		f[i] = make([]int, M)
	}
	for i := 1; i < M; i++ {
		for j := i - 1; j >= 1; j-- {
			f[j][i-j] = f[j-1][i-j] + f[j][i-j-1] - f[j-1][i-j-1] + cur*cur
			ans[cur] = f[j][i-j]
			cur++
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
	fmt.Fprintln(out, ans[n])
}
