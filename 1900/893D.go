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

	var n, d int
	fmt.Fscan(in, &n, &d)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	mn, mx := 0, 0
	ans := 0
	for _, x := range a {
		if x == 0 {
			if mn < 0 {
				mn = 0
			}
			if mx < 0 {
				mx = d
				ans++
			}
		} else {
			mn += x
			mx += x
			if mn > d {
				fmt.Fprintln(out, -1)
				return
			}
			if mx > d {
				mx = d
			}
		}
	}
	fmt.Fprintln(out, ans)
}
