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
	var a, b, c, d int
	fmt.Fscan(in, &a, &b, &c, &d)

	if n == 4 || k < 6+(n-5) {
		fmt.Fprintln(out, -1)
		return
	}

	nums := make([]int, 0, n-4)
	for i := 1; i <= n; i++ {
		if i != a && i != b && i != c && i != d {
			nums = append(nums, i)
		}
	}

	output := func(a, b, c, d int) {
		fmt.Fprint(out, a, c, " ")
		for _, i := range nums {
			fmt.Fprint(out, i, " ")
		}
		fmt.Fprintln(out, d, b)
	}

	output(a, b, c, d)
	output(c, d, a, b)
}
