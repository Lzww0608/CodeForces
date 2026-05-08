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

	a := []int{}
	for x := n; x > 0; x /= 3 {
		a = append(a, x%3)
	}

	m := len(a)
	for i := m - 1; i >= 0; i-- {
		if a[i] != 2 {
			continue
		}
		j := i + 1
		for j < m {
			if a[j] == 0 {
				a[j] = 1
				break
			}
			j++
		}

		if j == m {
			a = append(a, 1)
		}

		for k := range j {
			a[k] = 0
		}
		//fmt.Fprintln(out, a)

		ans := 0
		for k := len(a) - 1; k >= 0; k-- {
			ans = ans*3 + a[k]
		}
		fmt.Fprintln(out, ans)
		return
	}

	fmt.Fprintln(out, n)
}
