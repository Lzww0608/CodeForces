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

	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	mx, cnt := 0, 1

	st := []int{}
	m := make([]bool, n)
	for i := range n {
		if s[i] == '(' {
			st = append(st, i)
		} else {
			if len(st) > 0 {
				m[st[len(st)-1]] = true
				m[i] = true
				st = st[:len(st)-1]
			}
		}
	}

	cur := 0
	for _, x := range m {
		if x {
			cur++
			if cur > mx {
				mx, cnt = cur, 1
			} else if cur == mx {
				cnt++
			}
		} else {
			cur = 0
		}
	}

	fmt.Fprintln(out, mx, cnt)
}
