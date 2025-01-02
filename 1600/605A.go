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
	pos := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		pos[a[i]-1] = i
	}

	ans := 0
	for i := 0; i < n; {
		j := i + 1
		for j < n && pos[j] > pos[j-1] {
			j++
		}
		ans = max(ans, j-i)
		i = j
	}

	fmt.Fprintln(out, n-ans)
}
