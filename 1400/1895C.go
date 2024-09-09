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
	f := make([][]string, 6)
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		f[len(s)] = append(f[len(s)], s)
	}
	ans := 0
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if (i+j)&1 == 1 {
				continue
			}
			mid := (i + j) / 2
			cnt := map[int]int{}
			for _, s := range f[i] {
				cur := 0
				for k := range s {
					if k < mid {
						cur += int(s[k] - '0')
					} else {
						cur -= int(s[k] - '0')
					}
				}
				cnt[cur]++
			}
			for _, s := range f[j] {
				cur := 0
				for k := range s {
					if k+i < mid {
						cur += int(s[k] - '0')
					} else {
						cur -= int(s[k] - '0')
					}
				}
				ans += cnt[-cur]
			}
		}

	}
	fmt.Fprintln(out, ans)
}
