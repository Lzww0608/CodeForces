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
	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)

	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}

	p, o := 0, 0
	for _, x := range cnt {
		p += x / 2
		o += x & 1
	}

	ans := p / k * 2
	p %= k
	if o >= k-p*2 {
		ans++
	}

	fmt.Fprintln(out, ans)
}
