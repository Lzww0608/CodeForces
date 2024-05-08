package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t, n int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		var s string
		fmt.Fscan(in, &s)
		m := len(s)
		f := true
		for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
			if s[j] < s[i] {
				f = false
				break
			} else if s[j] > s[i] {
				break
			}
		}
		if f {
			if n&1 == 1 {
				fmt.Print(string(s))
			}
			fmt.Println(string(s))
		} else {
			if n&1 == 1 {
				for i := m - 1; i >= 0; i-- {
					fmt.Print(string(s[i]))
				}
				fmt.Println()
			} else {
				for i := m - 1; i >= 0; i-- {
					fmt.Print(string(s[i]))
				}
				fmt.Println(string(s))
			}
		}
	}
}
