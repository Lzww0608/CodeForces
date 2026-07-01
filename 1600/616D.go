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

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	mx, l, r := 0, 0, 0
	cnt := make(map[int]int)
	j := 0
	for i, x := range a {
		cnt[x]++
		for len(cnt) > k {
			if cnt[a[j]]--; cnt[a[j]] == 0 {
				delete(cnt, a[j])
			}
			j++
		}

		if i-j+1 > mx {
			mx = i - j + 1
			l, r = j+1, i+1
		}
	}

	fmt.Fprintln(out, l, r)
}
