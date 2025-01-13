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

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	cnt := make([]int, n+1)
	for i := range a {
		fmt.Fscan(in, &a[i])
		if a[i] < 0 {
			cnt[-a[i]]--
			m--
		} else {
			cnt[a[i]]++
		}
	}

	crime := 0
	for _, x := range cnt[1:] {
		if x == m {
			crime++
		}
	}

	for _, x := range a {
		if x > 0 {
			if cnt[x] != m {
				fmt.Fprintln(out, "Lie")
			} else if crime == 1 {
				fmt.Fprintln(out, "Truth")
			} else {
				fmt.Fprintln(out, "Not defined")
			}
		} else {
			if cnt[-x] != m {
				fmt.Fprintln(out, "Truth")
			} else if crime == 1 {
				fmt.Fprintln(out, "Lie")
			} else {
				fmt.Fprintln(out, "Not defined")
			}
		}

	}
}
