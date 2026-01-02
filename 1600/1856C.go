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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	ans := sort.Search(1e9, func(mx int) bool {
		for i := range a {
			cur, cnt := mx, k
			ok := false
			for _, x := range a[i:] {
				if x >= cur {
					ok = true
					break
				} else {
					cnt -= cur - x
					if cnt < 0 {
						break
					}
				}
				cur--
			}
			if ok && cnt >= 0 {
				return false
			}
		}

		return true
	})

	fmt.Fprintln(out, ans-1)
}
