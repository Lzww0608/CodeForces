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
	var a, b string
	fmt.Fscan(in, &n, &a, &b)
	if n&1 == 1 && a[n-1] != b[n-1] {
		fmt.Fprintln(out, "NO")
		return
	}
	cnt := 0
	for i := 0; i < n; i++ {
		if a[i] == '0' {
			cnt++
		}
		if b[i] == '0' {
			cnt--
		}
	}
	if cnt != 0 {
		fmt.Fprintln(out, "NO")
		return
	}
	x0, x1 := 0, 0
	same := 0
	for i := 0; i < n; i++ {
		if a[i] == '0' {
			x0++
		} else {
			x1++
		}

		if a[i] != b[i] {
			if same == 1 {
				fmt.Fprintln(out, "NO")
				return
			}
			same = -1
		} else {
			if same == -1 {
				fmt.Fprintln(out, "NO")
				return
			}
			same = 1
		}

		if x0 == x1 {
			same = 0
		}
	}

	fmt.Fprintln(out, "YES")
}
