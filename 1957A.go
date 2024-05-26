package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		m := map[int]int{}
		for i := range a {
			fmt.Fscan(in, &a[i])
			m[a[i]]++
		}
		ans := 0
		for _, v := range m {
			ans += v / 3
		}

		fmt.Fprintln(out, ans)

	}
}