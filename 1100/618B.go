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

	var n int
	fmt.Fscan(in, &n)
	a := make([][]int, n)
	b := make([][]int, n)
	for i := range a {
		b[i] = make([]int, n)
		a[i] = make([]int, n)
		for j := range a[i] {
			fmt.Fscan(in, &a[i][j])
			b[i][a[i][j]]++
		}
	}
	//fmt.Fprintln(out, b)
	ans := make([]int, n)
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			if b[j][i] == n-i {
				ans[j] = i
				break
			}
		}
	}
	for i := range ans {
		if ans[i] == 0 {
			ans[i] = n
			break
		}

	}

	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
	fmt.Fprintln(out)

}
