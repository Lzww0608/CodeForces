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
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	ans := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		ans += a[i]
	}
	b := make([]int, n)
	cnt := make(map[int]int)
	mx := 0

	sum := 0
	for i, x := range a {
		cnt[x]++
		if cnt[x] > 1 {
			mx = max(mx, x)
		}
		b[i] = mx
		sum += b[i]
	}
	ans += sum

	clear(cnt)
	mx, sum = 0, 0
	for i, x := range b {
		cnt[x]++
		if cnt[x] > 1 {
			mx = max(mx, x)
		}
		b[i] = mx
		sum += b[i]
	}
	ans += sum
	for i := 1; i < n; i++ {
		sum -= b[n-i]
		ans += sum
	}

	fmt.Fprintln(out, ans)

}
