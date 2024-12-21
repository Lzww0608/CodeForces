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
	var n, k int
	fmt.Fscan(in, &n, &k)
	if k%n == 0 {
		fmt.Fprintln(out, 0)
	} else {
		fmt.Fprintln(out, 2)
	}
	ans := make([][]int, n)
	j := 0
	p, q := k/n, k%n
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
		for t := 0; t < p; t++ {
			ans[i][j%n] = 1
			j++
		}
		if q > 0 {
			q--
			ans[i][j%n] = 1
			j++
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(out, ans[i][j])
		}
		fmt.Fprintln(out)
	}

}
