package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

const N int = 10

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
	var s string
	fmt.Fscan(in, &s)

	n := len(s)
	cnt := make([]int, N)
	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		cnt[x]++
	}

	ans := slices.Max(cnt)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}

			sum := 0
			for k := range n {
				x := int(s[k] - '0')
				if sum&1 == 0 {
					if x == i {
						sum++
					}
				} else {
					if x == j {
						sum++
					}
				}
			}

			ans = max(ans, sum-sum&1)
		}
	}

	fmt.Fprintln(out, n-ans)
}
