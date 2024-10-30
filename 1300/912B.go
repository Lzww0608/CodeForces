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

	var n, k int
	fmt.Fscan(in, &n, &k)
	t := 1
	for t <= n {
		t <<= 1
	}

	ans := 0
	if k == 1 {
		ans = n
	} else {
		ans = t - 1
	}

	fmt.Fprintln(out, ans)
}
