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

	var n, d, l int
	fmt.Fscan(in, &n, &d, &l)
	mx := l*(n+1)/2 - n/2
	mn := (n+1)/2 - n/2*l
	if d < mn || d > mx {
		fmt.Fprintln(out, -1)
		return
	}

	ans := make([]int, n)
	if d > 0 {
		for i := 1; i < n; i += 2 {
			ans[i] = 1
		}
		d += n / 2
		cnt := (n + 1) / 2
		for i := 0; i < n; i += 2 {
			cnt--
			t := min(l, d-cnt)
			ans[i] = t
			d -= t
		}
	} else if d < 0 {
		for i := 0; i < n; i += 2 {
			ans[i] = 1
		}
		d = -d + (n+1)/2
		cnt := n / 2
		for i := 1; i < n; i += 2 {
			cnt--
			t := min(l, d-cnt)
			ans[i] = t
			d -= t
		}
	} else {
		if n&1 == 0 {
			for i := 0; i < n; i++ {
				fmt.Fprint(out, 1, " ")
			}
		} else {
			fmt.Fprint(out, 1, 2, " ")
			for i := 2; i < n; i++ {
				fmt.Fprint(out, 1, " ")
			}
		}
		return
	}

	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}
