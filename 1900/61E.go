package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	v, pre, suf int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]Node, n)
	for i := range a {
		fmt.Fscan(in, &a[i].v)
	}
	ans := 0

	var mergeSort func(int, int)
	mergeSort = func(l, r int) {
		if l >= r {
			return
		}
		mid := l + ((r - l) >> 1)
		mergeSort(l, mid)
		mergeSort(mid+1, r)
		tmp := make([]Node, r-l+1)

		i, j, k := l, mid+1, 0
		for i <= mid || j <= r {
			if j > r || (i <= mid && a[i].v <= a[j].v) {
				ans += a[i].pre * (j - mid - 1)
				a[i].suf += j - mid - 1
				tmp[k] = a[i]
				i++
			} else {
				ans += a[j].suf * (mid + 1 - i)
				a[j].pre += mid + 1 - i
				tmp[k] = a[j]
				j++
			}
			k++
		}

		copy(a[l:r+1], tmp)
	}
	mergeSort(0, n-1)

	fmt.Fprintln(out, ans)
}
