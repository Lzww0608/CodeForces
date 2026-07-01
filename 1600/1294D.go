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

	var q, x int
	fmt.Fscan(in, &q, &x)
	cnt := make(map[int]int)
	cur := 0
	for range q {
		var y int
		fmt.Fscan(in, &y)
		y %= x
		cnt[y]++
		for cnt[cur%x] > 0 {
			cnt[cur%x]--
			cur++
		}
		fmt.Fprintln(out, cur)
	}
}
