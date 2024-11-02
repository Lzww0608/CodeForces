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
	var s string
	fmt.Fscan(in, &n, &s)
	cnt := make([]int, 26)
	for i := range s {
		cnt[int(s[i]&31)-1]++
	}
	ans := make([]byte, 0, n)
	for i, x := range cnt {
		for j := 0; j < x; j++ {
			ans = append(ans, byte('a'+i))
		}
	}
	fmt.Fprintln(out, string(ans))
}
