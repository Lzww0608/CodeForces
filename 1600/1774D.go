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
	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([][]int, n)
	cnt := make([]int, n)
	tot := 0
	for i := range a {
		a[i] = make([]int, m)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
			if a[i][j] == 1 {
				tot++
				cnt[i]++
			}
		}
	}

	if tot%n != 0 {
		fmt.Fprintln(out, -1)
		return
	}
	avg := tot / n

	ans := [][]int{}
	for j := 0; j < m; j++ {
		more := []int{}
		less := []int{}
		for i := 0; i < n; i++ {
			if a[i][j] == 1 && cnt[i] > avg {
				more = append(more, i)
			} else if a[i][j] == 0 && cnt[i] < avg {
				less = append(less, i)
			}
		}
		for i := 0; i < min(len(more), len(less)); i++ {
			ans = append(ans, []int{more[i], less[i], j})
			cnt[more[i]]--
			cnt[less[i]]++
		}
	}

	fmt.Fprintln(out, len(ans))
	for _, v := range ans {
		fmt.Fprintln(out, v[0]+1, v[1]+1, v[2]+1)
	}
}
