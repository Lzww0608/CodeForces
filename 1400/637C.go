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
	str := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &str[i])
	}

	ans := 6
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			cnt := 0
			for k := 0; k < 6; k++ {
				if str[i][k] != str[j][k] {
					cnt++
				}
			}
			ans = min(ans, (cnt+1)/2-1)
		}
	}

	fmt.Fprintln(out, max(ans, 0))
}
