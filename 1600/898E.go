package main

import (
	"bufio"
	"fmt"
	"math"
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
	b := make([]int, 0, n)
	c := make([]int, 0, n)
	zero := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		x := int(math.Sqrt(float64(a[i])))
		if x*x != a[i] {
			b = append(b, a[i])
		} else {
			c = append(c, a[i])
			if a[i] == 0 {
				zero++
			}
		}
	}
	sort.Ints(b)

	ans := 0
	cnt := len(c)
	if cnt == n/2 {
		fmt.Fprintln(out, 0)
		return
	} else if cnt > n/2 {
		ans += cnt - n/2
		if zero > n/2 {
			ans += zero - n/2
		}
	} else {
		cnt = n/2 - cnt
		for i := range b {
			x := int(math.Sqrt(float64(b[i])))
			t := min(b[i]-x*x, (x+1)*(x+1)-b[i])
			b[i] = t
		}
		sort.Ints(b)
		for i := 0; i < cnt; i++ {
			ans += b[i]
		}
	}

	fmt.Fprintln(out, ans)
}
