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
	var n, m, k int
	fmt.Fscan(in, &n, &m, &k)
	a := make([]int, k)
	b := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &b[i])
	}

	cnt1 := make(map[int]int)
	cnt2 := make(map[int]int)
	cnt3 := make(map[[2]int]int)
	ans := 0
	for i := 0; i < k; i++ {
		ans += i - cnt1[a[i]] - cnt2[b[i]] + cnt3[[2]int{a[i], b[i]}]
		cnt1[a[i]]++
		cnt2[b[i]]++
		cnt3[[2]int{a[i], b[i]}]++
	}

	fmt.Fprintln(out, ans)
}
