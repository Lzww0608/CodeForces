package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

const N int = 100_001

var F [N][]int

func init() {
	for i := 2; i < N; i++ {
		for j := i; j < N; j += i {
			F[j] = append(F[j], i)
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
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	ans := math.MaxInt32
	freq := make([]int, m+1)
	cnt := m - 1
	for l, r := 0, 0; r < n; r++ {
		for _, x := range F[a[r]] {
			if x > m {
				break
			}
			if freq[x]++; freq[x] == 1 {
				cnt--
			}
		}

		for cnt == 0 && l <= r {
			ans = min(ans, a[r]-a[l])
			for _, x := range F[a[l]] {
				if x > m {
					break
				}
				if freq[x]--; freq[x] == 0 {
					cnt++
				}
			}

			l++
		}
	}

	if ans == math.MaxInt32 {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, ans)
	}
}
