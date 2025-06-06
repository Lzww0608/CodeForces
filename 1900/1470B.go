package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1_000_005

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	a := make([]int, N)
	for i := 2; i < N; i++ {
		if a[i] == 0 {
			for j := i; j < N; j += i {
				if a[j] == 0 {
					a[j] = i
				}
			}
		}
	}

	core := func(x int) int {
		cnt := 1
		for x > 1 {
			p := a[x]
			e := 1
			for x /= p; a[x] == p; x /= p {
				e ^= 1
			}
			if e > 0 {
				cnt *= p
			}
		}

		return cnt
	}

	var t, n, v, q, w int
	for fmt.Fscan(in, &t); t > 0; t-- {
		cnt := make(map[int]int)
		for fmt.Fscan(in, &n); n > 0; n-- {
			fmt.Fscan(in, &v)
			cnt[core(v)]++
		}
		maxC, c1 := 0, cnt[1]
		for v, c := range cnt {
			maxC = max(maxC, c)
			if v > 1 && (c&1) == 0 {
				c1 += c
			}
		}

		for fmt.Fscan(in, &q); q > 0; q-- {
			if fmt.Fscan(in, &w); w == 0 {
				fmt.Fprintln(out, maxC)
			} else {
				fmt.Fprintln(out, max(maxC, c1))
			}
		}

	}
}
