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
	var n, m int
	fmt.Fscan(in, &n, &m)
	a, b := 0, 0
	for t := n; t%2 == 0; t /= 2 {
		a++
	}
	for t := n; t%5 == 0; t /= 5 {
		b++
	}

	x := 1
	for x <= m {
		if a > b {
			if x*5 <= m {
				x *= 5
			} else {
				break
			}
			b++
		} else if a < b {
			if x*2 <= m {
				x *= 2
			} else {
				break
			}
			a++
		} else {
			if x*10 <= m {
				x *= 10
			} else {
				break
			}
		}
	}

	x = m / x * x

	fmt.Fprintln(out, n*x)
}
