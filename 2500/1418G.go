package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

const N int = 3

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	base := make([]uint64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		a[i]--
		base[i] = rand.Uint64()
	}

	cnt := make([]int, n)
	h := make([]uint64, n+1)
	cntH := map[uint64]int{0: 1}
	rem := make([]int, n)
	l := 0

	ans := 0
	for i, x := range a {
		cnt[x]++
		for cnt[x] > N {
			cnt[a[l]]--
			cntH[h[l]]--
			l++
		}
		h[i+1] = h[i] + uint64((rem[x]+1)%N-(rem[x])%N)*base[x]
		ans += cntH[h[i+1]]
		cntH[h[i+1]]++
		rem[x]++
	}

	fmt.Fprint(out, ans)
}
