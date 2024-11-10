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

	var n, q, x int
	fmt.Fscan(in, &n)
	a := map[int]int{}
	four, two := 0, 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x)
		a[x]++
		if a[x]%4 == 0 {
			four++
			two--
		} else if a[x]%2 == 0 {
			two++
		}
	}

	fmt.Fscan(in, &q)
	var s, t string
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &s, &t)
		x, _ = strconv.Atoi(t)
		//fmt.Fprintln(out, x)
		if s == "-" {
			if a[x]%4 == 0 {
				four--
				two++
			} else if a[x]%4 == 2 {
				two--
			}
			a[x]--
		} else {
			if a[x]%4 == 3 {
				two--
				four++
			} else if a[x]%4 == 1 {
				two++
			}
			a[x]++
		}
		//fmt.Fprintln(out, four, two)
		if four >= 2 || (four >= 1 && two >= 2) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
