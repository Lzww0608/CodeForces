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
	s := make([][]byte, n)
	cnt := [3]int{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s[i])
		for j := 0; j < n; j++ {
			if s[i][j] == 'X' {
				cnt[(i+j)%3]++
			}
		}
	}

	mn := 0
	for i := range cnt {
		if cnt[i] < cnt[mn] {
			mn = i
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if s[i][j] == 'X' && (i+j)%3 == mn {
				fmt.Fprintf(out, "%c", 'O')
			} else {
				fmt.Fprintf(out, "%c", s[i][j])
			}
		}
		fmt.Fprintln(out)
	}

}
