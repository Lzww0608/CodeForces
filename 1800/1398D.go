package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var R, G, B int
	fmt.Fscan(in, &R, &G, &B)
	r := make([]int, R)
	g := make([]int, G)
	b := make([]int, B)
	for i := range r {
		fmt.Fscan(in, &r[i])
	}
	for i := range g {
		fmt.Fscan(in, &g[i])
	}
	for i := range b {
		fmt.Fscan(in, &b[i])
	}

	sort.Slice(r, func(i, j int) bool {
		return r[i] > r[j]
	})
	sort.Slice(g, func(i, j int) bool {
		return g[i] > g[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})

	dp := make([][][]int, R+1)
	for i := range dp {
		dp[i] = make([][]int, G+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, B+1)
		}
	}

	ans := 0
	for i := 0; i <= R; i++ {
		for j := 0; j <= G; j++ {
			for k := 0; k <= B; k++ {
				cur := dp[i][j][k]
				if cur > ans {
					ans = cur
				}
				if i < R && j < G {
					next := cur + r[i]*g[j]
					if next > dp[i+1][j+1][k] {
						dp[i+1][j+1][k] = next
					}
				}
				if i < R && k < B {
					next := cur + r[i]*b[k]
					if next > dp[i+1][j][k+1] {
						dp[i+1][j][k+1] = next
					}
				}
				if j < G && k < B {
					next := cur + g[j]*b[k]
					if next > dp[i][j+1][k+1] {
						dp[i][j+1][k+1] = next
					}
				}
			}
		}
	}

	fmt.Fprintln(out, ans)
}
