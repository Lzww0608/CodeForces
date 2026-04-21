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
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	cnt := make(map[int]int)
	ans := 0
	for _, x := range a {
		for y := 1; y < (1 << 30); y <<= 1 {
			ans += cnt[y-x]
		}
		cnt[x]++

	}

	fmt.Fprintln(out, ans)
}
