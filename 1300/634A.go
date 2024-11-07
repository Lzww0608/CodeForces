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
	b := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
	idA, idB := 0, 0
	for i := range b {
		fmt.Fscan(in, &b[i])
		if b[i] == 0 {
			idB = (i + 1) % n
		}
	}
	for i := 0; i < n; i++ {
		if a[i] == b[idB] {
			idA = i
			break
		}
	}
	for i := 0; i < n; i++ {
		x := (idA + i) % n
		if a[x] == 0 {
			continue
		}

		if a[x] != b[idB] {
			fmt.Fprintln(out, "NO")
			return
		}
		idB = (idB + 1) % n
	}
	fmt.Fprintln(out, "YES")
}
