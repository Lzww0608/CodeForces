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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := 1
	i := 0
	for i < n && a[i] == 0 {
		i++
	}

	if i == n {
		fmt.Fprintln(out, 0)
		return
	}

	pre := 1
	for i < n {
		if a[i] == 1 {
			ans *= pre
			pre = 1
		} else {
			pre++
		}
		i++
	}

	fmt.Fprintln(out, ans)
}
