package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const MOD int = 1_000_000_007
const N int = 200_001

var F [N]int
var invF [N]int

func init() {
	F[0], invF[0] = 1, 1
	for i := 1; i < N; i++ {
		F[i] = (F[i-1] * i) % MOD
	}
	invF[N-1] = quickPow(F[N-1], MOD-2)
	for i := N - 2; i >= 1; i-- {
		invF[i] = invF[i+1] * (i + 1) % MOD
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
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	ans := 0
	for i, x := range a {
		j := sort.SearchInts(a, x+k+1)
		cnt := j - i - 1
		if cnt < m-1 {
			continue
		}
		ans = (ans + F[cnt]*invF[m-1]%MOD*invF[cnt-(m-1)]%MOD) % MOD
	}
	fmt.Fprintln(out, ans)
}

func quickPow(a, r int) int {
	res := 1
	for r > 0 {
		if r&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		r >>= 1
	}
	return res
}
