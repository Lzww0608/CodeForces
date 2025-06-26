package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	val    int
	needOR bool
}
type segmentTree []node

func (t segmentTree) _pushUp(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	if lo.needOR {
		t[o].val = lo.val | ro.val
	} else {
		t[o].val = lo.val ^ ro.val
	}
}

func (t segmentTree) _build(arr []int, o, l, r int) {
	if l == r {
		t[o].val = arr[l]
		t[o].needOR = true
		return
	}
	mid := (l + r) >> 1
	t._build(arr, o<<1, l, mid)
	t._build(arr, o<<1|1, mid+1, r)
	t[o].needOR = !t[o<<1].needOR
	t._pushUp(o)
}

func (t segmentTree) _update(o, l, r int, idx int, val int) {
	if l == r {
		t[o].val = val
		return
	}
	mid := (l + r) >> 1
	if idx <= mid {
		t._update(o<<1, l, mid, idx, val)
	} else {
		t._update(o<<1|1, mid+1, r, idx, val)
	}
	t._pushUp(o)
}

func (t segmentTree) init(arr []int)                 { t._build(arr, 1, 1, len(arr)-1) }
func (t segmentTree) update(n int, idx int, val int) { t._update(1, 1, n, idx, val) }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	n = 1 << uint(n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	t := make(segmentTree, n*2)
	t.init(a)
	for ; m > 0; m-- {
		var idx, val int
		fmt.Fscan(in, &idx, &val)
		t.update(n, idx, val)
		fmt.Fprintln(out, t[1].val)
	}
}
