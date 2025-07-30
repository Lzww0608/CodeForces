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

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	cnt := make(map[int]int)
	for i := range a {
		fmt.Fscan(in, &a[i])
		for j := 2; j*j <= a[i]; j++ {
			for a[i]%j == 0 {
				cnt[j]++
				a[i] /= j
			}
		}
		if a[i] != 1 {
			cnt[a[i]]++
		}
	}

	for _, x := range cnt {
		if x%n != 0 {
			fmt.Fprintln(out, "NO")
			return
		}
	}
	fmt.Fprintln(out, "YES")
	return
}
