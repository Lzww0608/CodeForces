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
	ans := 0
	for i := 0; i < n; i++ {
		var x, m int
		fmt.Fscan(in, &x, &m)
		ans ^= solve(x-1) ^ solve(m+x-1)
	}
	if ans == 0 {
		fmt.Fprintln(out, "bolik")
	} else {
		fmt.Fprintln(out, "tolik")
	}
}

func solve(x int) int {
	switch x % 4 {
	case 0:
		return x
	case 1:
		return 1
	case 2:
		return x + 1
	case 3:
		return 0
	}

	return 0
}
