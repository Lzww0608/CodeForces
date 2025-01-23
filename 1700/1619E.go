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
		cnt[a[i]]++
	}

	ans := 0
	s := []int{}
	for i := 0; i <= n; i++ {
		if ans == -1 {
			fmt.Fprint(out, -1, " ")
			continue
		}
		fmt.Fprint(out, ans+cnt[i], " ")
		for cnt[i] > 0 {
			cnt[i]--
			s = append(s, i)
		}

		if len(s) == 0 {
			ans = -1
		} else {
			ans += i - s[len(s)-1]
			s = s[:len(s)-1]
		}
	}

	fmt.Fprintln(out)
}
