package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	sort.Ints(a)
	ans := 0
	pre := 0
	for _, x := range a {
		if x > pre+1 {
			pre = x - 1
			ans++
		} else if x == pre+1 {
			pre = x
			ans++
		} else if x == pre {
			pre = x + 1
			ans++
		}
	}

	fmt.Fprintln(out, ans)
}
