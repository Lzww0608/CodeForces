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

	var q int
	for fmt.Fscan(in, &q); q > 0; q-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var s string
	fmt.Fscan(in, &s)
	cnt := make([]int, 4)
	for i := range s {
		switch s[i] {
		case 'U':
			cnt[0]++
		case 'D':
			cnt[1]++
		case 'L':
			cnt[2]++
		case 'R':
			cnt[3]++
		}
	}

	mn_x := min(cnt[2], cnt[3])
	mn_y := min(cnt[0], cnt[1])
	if mn_x == 0 && mn_y == 0 {
		fmt.Fprintln(out, 0)
		fmt.Fprintln(out)
		return
	} else if mn_x == 0 {
		fmt.Fprintln(out, 2)
		fmt.Fprintln(out, "UD")
		return
	} else if mn_y == 0 {
		fmt.Fprintln(out, 2)
		fmt.Fprintln(out, "LR")
		return
	}
	fmt.Fprintln(out, (mn_x+mn_y)*2)
	for i := 0; i < mn_x; i++ {
		fmt.Fprint(out, "L")
	}
	for i := 0; i < mn_y; i++ {
		fmt.Fprint(out, "U")
	}
	for i := 0; i < mn_x; i++ {
		fmt.Fprint(out, "R")
	}
	for i := 0; i < mn_y; i++ {
		fmt.Fprint(out, "D")
	}
	fmt.Fprintln(out)
}
