package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	cnt := make([]int, 26)
	m := make([]int, 26*26)
	for i := range s {
		x := int(s[i] - 'a')
		for j := range 26 {
			m[x*26+j] += cnt[j]
		}

		cnt[x]++
	}

	ans := max(slices.Max(m), slices.Max(cnt))
	fmt.Fprintln(out, ans)
}
