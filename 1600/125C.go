package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1_000_001

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	k := 1
	for k*(k+1)/2 <= n {
		k++
	}
	ans := make([][]int, k)
	t := 0
	for i := 0; i < k; i++ {
		for j := i + 1; j < k; j++ {
			ans[i] = append(ans[i], t)
			ans[j] = append(ans[j], t)
			t++
		}
	}
	fmt.Fprintln(out, len(ans))
	for i := 0; i < k; i++ {
		for _, v := range ans[i] {
			fmt.Fprint(out, v+1, " ")
		}
		fmt.Fprintln(out)
	}
}
