package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	var s string
	vis := make(map[string]bool)
	cnt := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		if vis[s] {
			cnt[s]++
			x := cnt[s]
			t := s + strconv.Itoa(x)
			vis[t] = true
			fmt.Fprintln(out, t)
		} else {
			vis[s] = true
			fmt.Fprintln(out, "OK")
		}
	}
}
