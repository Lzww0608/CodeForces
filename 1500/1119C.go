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
	s := make([][]int, n)
	t := make([][]int, n)

	row := make([]int, n)
	col := make([]int, m)
	for i := 0; i < n; i++ {
		s[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &s[i][j])
			if s[i][j] == 0 {
				row[i] ^= 1
				col[j] ^= 1
			}
		}
	}

	for i := 0; i < n; i++ {
		t[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &t[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if t[i][j] == 0 {
				row[i] ^= 1
				col[j] ^= 1
			}
		}
		if row[i] != 0 {
			fmt.Fprintln(out, "No")
			return
		}
	}

	for i := 0; i < m; i++ {
		if col[i] != 0 {
			fmt.Fprintln(out, "No")
			return
		}
	}

	fmt.Fprintln(out, "Yes")
}
