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
	cntA := make([]int, 10)
	cntB := make([]int, 10)
	var s string
	fmt.Fscan(in, &s)
	for i := 0; i < n; i++ {
		cntA[s[i]-'0']++
	}
	fmt.Fscan(in, &s)
	for i := 0; i < n; i++ {
		cntB[s[i]-'0']++
	}

	cur, ans := 0, 0
	for i := 0; i < 10; i++ {
		cur += cntA[i]
		t := min(cur, cntB[i])
		cur -= t
		ans += t
	}
	fmt.Fprintln(out, n-ans)

	cur, ans = cntB[9], 0
	for i := 8; i >= 0; i-- {
		t := min(cur, cntA[i])
		cur -= t
		ans += t
		cur += cntB[i]
	}
	fmt.Fprintln(out, ans)
}
