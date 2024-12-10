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
	cur, cnt := k, 0
	exist := make(map[int]bool)
	mx := 0
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &a[i])
		exist[a[i]] = true
		cnt++
		if a[i] == cur {
			for exist[cur] {
				cnt--
				cur--
			}
		}
		mx = max(mx, cnt)
	}
	if mx > n*m-4 {
		fmt.Fprintln(out, "TIDAK")
	} else {
		fmt.Fprintln(out, "YA")
	}
}
