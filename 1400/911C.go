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

	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)
	cnt := map[int]int{}
	cnt[a]++
	cnt[b]++
	cnt[c]++
	if cnt[1] >= 1 {
		fmt.Fprintln(out, "YES")
		return
	} else if cnt[2] >= 2 {
		fmt.Fprintln(out, "YES")
		return
	} else if cnt[3] == 3 {
		fmt.Fprintln(out, "YES")
		return
	} else if cnt[4] == 2 && cnt[2] == 1 {
		fmt.Fprintln(out, "YES")
		return
	}
	fmt.Fprintln(out, "NO")
}
