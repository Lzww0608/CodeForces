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

	var k int
	fmt.Fscan(in, &k)
	mask := 10
	for mask <= k {
		mask *= 10
	}
	mask /= 10
	cnt := 0
	ans := []int{}
	for t := mask; t > 0; t /= 10 {
		cnt = max(cnt, (k/t)%10)
	}
	fmt.Fprintln(out, cnt)
	for k > 0 {
		for k >= mask {
			cur := 0
			for t := mask; t > 0; t /= 10 {
				if (k/t)%10 != 0 {
					cur += t
					k -= t
				}
			}
			ans = append(ans, cur)
		}
		mask /= 10
	}
	for _, x := range ans {
		fmt.Fprint(out, x, " ")
	}
}
