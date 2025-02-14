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
	if m < n-1 {
		fmt.Fprintln(out, "Impossible")
		return
	}

	ans := make([][2]int, m)
	cur := 0
	for i := 0; i < n && cur < m; i++ {
		for j := i + 1; j < n && cur < m; j++ {
			if gcd(i+1, j+1) == 1 {
				ans[cur] = [2]int{i + 1, j + 1}
				cur++
			}
		}
	}

	if cur != m {
		fmt.Fprintln(out, "Impossible")
		return
	}

	fmt.Fprintln(out, "Possible")
	for _, v := range ans {
		fmt.Fprintln(out, v[0], v[1])
	}
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}
