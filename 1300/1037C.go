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
	var n int
	fmt.Fscan(in, &n)
	var a, b string
	fmt.Fscan(in, &a)
	fmt.Fscan(in, &b)

	pre := -2
	ans := 0
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			if i-pre > 1 || a[i-1] != b[i] {
				ans++
				pre = i
			} else {
				pre = -2
			}

		}
	}

	fmt.Fprintln(out, ans)
}
