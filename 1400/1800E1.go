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
	var n, k int
	fmt.Fscan(in, &n, &k)
	var s, t string
	fmt.Fscan(in, &s, &t)
	cnt := make([]int, 26)
	for i := 0; i < n; i++ {
		if i < k && i+k >= n {
			if s[i] != t[i] {
				fmt.Fprintln(out, "NO")
				return
			}
		} else {
			cnt[int(s[i]&31)-1]++
			cnt[int(t[i]&31)-1]--
		}
	}
	for _, x := range cnt {
		if x != 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
	return
}
