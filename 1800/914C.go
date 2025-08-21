package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const N int = 1001

const MOD int = 1_000_000_007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	var k int
	fmt.Fscan(in, &k)

	if k == 0 {
		fmt.Fprintln(out, 1)
		return
	}

	f := make([]int, N)
	for i := 2; i < N; i++ {
		j := bits.OnesCount(uint(i))
		f[i] = f[j] + 1
	}

	C := [N][N]int{}
	for i := 0; i < N; i++ {
		C[i][0] = 1
		for j := 1; j <= i; j++ {
			C[i][j] = (C[i-1][j-1] + C[i-1][j]) % MOD
		}
	}

	ans := 0
	n := len(s)
	for i := 1; i <= n; i++ {
		if f[i] != k-1 {
			continue
		}

		cnt := i
		for j, c := range s {
			if c == '1' {
				ans = (ans + C[n-j-1][cnt]) % MOD
				cnt--
				if cnt < 0 {
					break
				}
			}
		}
	}

	cnt := 0
	for i := range s {
		if s[i] == '1' {
			cnt++
		}
	}
	if f[cnt] == k-1 {
		ans = (ans + 1) % MOD
	}

	if k == 1 {
		ans = (ans - 1 + MOD) % MOD
	}

	fmt.Fprintln(out, ans)
}
