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
	str := make([]string, n)
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = append(ans[i], i+1)
		fmt.Fscan(in, &str[i])
		for j := 0; j < n; j++ {
			if str[i][j] == '1' {
				ans[j] = append(ans[j], i+1)
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprint(out, len(ans[i]), " ")
		for _, x := range ans[i] {
			fmt.Fprint(out, x, " ")
		}
		fmt.Fprintln(out)
	}
}
