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
	var s, t string
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &s, &t)
	posA := []int{}
	posB := []int{}
	for i := 0; i < n; i++ {
		if s[i] != t[i] {
			if s[i] == 'a' {
				posA = append(posA, i+1)
			} else {
				posB = append(posB, i+1)
			}
		}
	}

	if (len(posA)+len(posB))&1 == 1 {
		fmt.Fprintln(out, -1)
		return
	} else if len(posA)+len(posB) == 0 {
		fmt.Fprintln(out, 0)
		return
	}

	ans := len(posA)/2 + len(posB)/2 + 2*(len(posA)&1)
	fmt.Fprintln(out, ans)
	for i := 0; i < len(posA)-1; i += 2 {
		fmt.Fprintln(out, posA[i], posA[i+1])
	}
	for i := 0; i < len(posB)-1; i += 2 {
		fmt.Fprintln(out, posB[i], posB[i+1])
	}
	if len(posA)&1 == 1 {
		fmt.Fprintln(out, posA[len(posA)-1], posA[len(posA)-1])
		fmt.Fprintln(out, posB[len(posB)-1], posA[len(posA)-1])
	}
}
