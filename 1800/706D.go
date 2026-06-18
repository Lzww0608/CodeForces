package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 200_000 * 30

var trie [N][2]int
var sum [N]int
var cnt int = 1

func add(x, op int) {
	p := 1
	sum[p] += op
	for i := 29; i >= 0; i-- {
		y := (x >> i) & 1
		if trie[p][y] == 0 {
			cnt++
			trie[p][y] = cnt
		}
		p = trie[p][y]
		sum[p] += op
	}
}

func query(x int) (res int) {
	p := 1
	for i := 29; i >= 0; i-- {
		y := (x >> i) & 1
		if trie[p][y^1] != 0 && sum[trie[p][y^1]] != 0 {
			res |= 1 << i
			p = trie[p][y^1]
		} else {
			p = trie[p][y]
		}
	}

	return
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	cnt = 1
	var q int
	fmt.Fscan(in, &q)
	var s string
	var x int
	for range q {
		fmt.Fscan(in, &s, &x)
		if s == "+" {
			add(x, 1)
		} else if s == "-" {
			add(x, -1)
		} else {
			fmt.Fprintln(out, max(x, query(x)))
		}
	}
}
