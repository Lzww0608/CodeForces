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

	var n, a, b int
	fmt.Fscan(in, &n, &a, &b)
	pre, sum := 1, 1
	ans := make([]int, 0, n)
	ans = append(ans, 1)
	for i := 1; i < n; i++ {
		if b > 0 {
			b--
			x := sum + 1
			ans = append(ans, x)
			pre = x
			sum += x
		} else if a > 0 {
			if pre+1 > sum {
				ans = append(ans, 1)
				sum += 1
				pre = 1
				continue
			}
			a--
			x := pre + 1
			ans = append(ans, x)
			pre = x
			sum += x
		} else {
			ans = append(ans, 1)
		}
		if ans[len(ans)-1] > 50_000 {
			fmt.Fprintln(out, -1)
			return
		}
	}

	if a > 0 || b > 0 {
		fmt.Fprintln(out, -1)
		return
	}

	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}
