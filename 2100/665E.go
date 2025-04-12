package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	cnt      int
	children [2]*Node
}

func (n *Node) insert(x int) {
	cur := n
	for i := 29; i >= 0; i-- {
		d := (x >> i) & 1
		if cur.children[d] == nil {
			cur.children[d] = new(Node)
		}
		cur = cur.children[d]
		cur.cnt++
	}
	return
}

func (n *Node) countLimitXor(v, k int) (cnt int) {
	cur := n
	for i := 29; i >= 0 && cur != nil; i-- {
		b := (v >> i) & 1
		if (k>>i)&1 == 1 {
			if cur.children[b] != nil {
				cnt += cur.children[b].cnt
			}
			b ^= 1
		}
		cur = cur.children[b]
	}
	return
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	ans := n * (n + 1) / 2
	t := &Node{}
	t.insert(0)
	sum := 0
	for _, x := range a {
		sum ^= x
		ans -= t.countLimitXor(sum, k)
		t.insert(sum)
	}

	fmt.Fprintln(out, ans)
}
