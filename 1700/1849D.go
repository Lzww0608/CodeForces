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
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	ans := 0
	for i := 0; i < n; {
		if a[i] == 0 {
			i++
			ans++
			continue
		}
		l := i
		two := false
		for i < n && a[i] != 0 {
			if a[i] == 2 {
				two = true
			}
			i++
		}

		ans++
		if l > 0 && a[l-1] == 0 {
			ans--
			if !two {
				continue
			}
		}
		if i < n {
			a[i] = 1
			i++
		}
	}
	fmt.Fprintln(out, ans)
}
