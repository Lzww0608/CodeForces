package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

type Trie struct {
	next [N]*Trie
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s, t string
	fmt.Fscan(in, &s, &t)
	n := len(s)
	var k int
	fmt.Fscan(in, &k)

	ans := 0
	root := &Trie{}

	for i := range n {
		cnt := 0
		cur := root
		for j := i; j < n; j++ {
			x := int(s[j] - 'a')
			if t[x] == '0' {
				cnt++
			}

			if cnt > k {
				break
			}

			if cur.next[x] == nil {
				cur.next[x] = &Trie{}
				ans++
			}
			cur = cur.next[x]
		}
	}

	fmt.Fprintln(out, ans)
}
