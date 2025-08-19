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
	var k, x int
	fmt.Fscan(in, &k, &x)

	if x >= k*k {
		fmt.Fprintln(out, k*2-1)
		return
	}

	if k*(k+1)/2 >= x {
		l, r := 1, k
		ans := 0
		for l <= r {
			mid := l + ((r - l) >> 1)
			t := mid * (mid + 1) / 2
			if t >= x {
				ans = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		fmt.Fprintln(out, ans)
	} else {
		x -= k * (k + 1) / 2

		l, r := 1, k-1
		ans := k - 1
		sum := k * (k - 1) / 2
		for l <= r {
			mid := l + ((r - l) >> 1)
			t := sum - (k-1-mid)*(k-mid)/2
			if t < x {
				l = mid + 1
			} else {
				ans = mid
				r = mid - 1
			}
		}
		fmt.Fprintln(out, k+ans)
	}
}
