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
	var s string
	fmt.Fscan(in, &s)
	m := make(map[byte]int)
	k := 0

	for i := range s {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = k
			k++
		}
	}

	pos := make([][]int, k)
	for i := range s {
		j := m[s[i]]
		pos[j] = append(pos[j], i)
	}

	ans := n
	mx := 0
	for i := range n {
		for j := range k {
			mx = max(mx, pos[j][0])
		}
		ans = min(ans, mx-i+1)

		j := m[s[i]]
		pos[j] = pos[j][1:]
		if len(pos[j]) == 0 {
			break
		}
	}

	fmt.Fprintln(out, ans)
}
