package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

type Trie struct {
	next  [N]*Trie
	isEnd bool
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
	insert := func(a string) {
		cur := root
		for i := 0; i < len(a); i++ {
			c := a[i] - 'a'
			if cur.next[c] == nil {
				cur.next[c] = &Trie{}
			}
			cur = cur.next[c]
		}

		if !cur.isEnd {
			ans++
		}
		cur.isEnd = true

	}

	for i := range n {
		cnt := 0
		for j := i; j < n; j++ {
			x := int(s[j] - 'a')
			if t[x] == '0' {
				cnt++
			}

			if cnt > k {
				break
			}

			insert(s[i : j+1])
		}
	}

	fmt.Fprintln(out, ans)
}
