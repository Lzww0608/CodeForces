package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	sort.Ints(a)

	l, r := 0, n
	ans := 0
	for l < r {
		mid := l + ((r - l) >> 1)
		f := true
		if mid*2 >= n {
			f = false
		} else {
			b := make([]int, n)
			i, j := 0, n-(mid+1)
			for k := 0; k < mid*2+1; k++ {
				if k&1 == 0 {
					b[k] = a[j]
					j++
				} else {
					b[k] = a[i]
					i++
				}
			}

			for i := 1; i < mid*2+1; i += 2 {
				if b[i] >= b[i-1] || b[i] >= b[i+1] {
					f = false
					break
				}
			}

		}
		if f {
			ans = mid
			l = mid + 1
		} else {
			r = mid
		}
	}

	fmt.Fprintln(out, ans)
	b := make([]int, n)
	i, j := 0, n-(ans+1)
	for k := 0; k < ans*2+1; k++ {
		if k&1 == 0 {
			b[k] = a[j]
			j++
		} else {
			b[k] = a[i]
			i++
		}
	}

	for k := ans*2 + 1; k < n; k++ {
		b[k] = a[ans+(k-ans*2-1)]
	}
	for _, x := range b {
		fmt.Fprint(out, x, " ")
	}
}
