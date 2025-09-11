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
	cnt := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
		cnt[a[i]]++
	}

	b := true
	for i := 0; i < n; i++ {
		if cnt[i] == 0 {
			fmt.Fprintln(out, i)
			return
		} else if cnt[i] == 1 {
			if b {
				b = false
			} else {
				fmt.Fprintln(out, i)
				return
			}
		}
	}

	fmt.Fprintln(out, n)
}
