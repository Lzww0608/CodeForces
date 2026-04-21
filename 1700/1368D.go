package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 20

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

	cnt := [N]int{}
	for i := range N {
		for _, x := range a {
			if x&(1<<i) != 0 {
				cnt[i]++
			}
		}
	}

	ans := 0
	for range n {
		x := 0
		for j := range cnt {
			if cnt[j] > 0 {
				cnt[j]--
				x |= 1 << j
			}

		}
		ans += x * x
	}

	fmt.Fprintln(out, ans)
}
