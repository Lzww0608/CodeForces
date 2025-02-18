package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1_005

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, x int
	cnt := make([]int, N)
	fmt.Fscan(in, &n)
	for i := 0; i < n*n; i++ {
		fmt.Fscan(in, &x)
		cnt[x]++
	}

	a := []pair{}
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	for i := 0; i < (n+1)/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			if i != n-i-1 && j != n-j-1 {
				a = append(a, pair{4, i, j})
			} else if i == n-i-1 && j == n-j-1 {
				a = append(a, pair{1, i, j})
			} else {
				a = append(a, pair{2, i, j})
			}
		}
	}

	for _, d := range []int{4, 2, 1} {
		k := 1
		for _, v := range a {
			if v.d != d {
				continue
			}
			i, j := v.x, v.y
			for k < N && cnt[k] < d {
				k++
			}
			if k == N {
				fmt.Fprintln(out, "NO")
				return
			}
			ans[i][j], ans[i][n-j-1], ans[n-i-1][j], ans[n-i-1][n-j-1] = k, k, k, k
			cnt[k] -= d
		}
	}

	fmt.Fprintln(out, "YES")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(out, ans[i][j], " ")
		}
		fmt.Fprintln(out, "")
	}
}

type pair struct {
	d, x, y int
}
