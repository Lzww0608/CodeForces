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
 
	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}
 
func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(in, &a[i])
	}
 
	sort.Ints(a)
	if k >= 3 {
		fmt.Fprintln(out, 0)
		return
	} else if k == 1 {
		ans := a[0]
		for i := 1; i < n; i++ {
			ans = min(ans, a[i]-a[i-1])
		}
		fmt.Fprintln(out, ans)
		return
	}
 
	ans := a[0]
	for i := 1; i < n; i++ {
		ans = min(ans, a[i]-a[i-1])
		x := a[i]
		j := i
		for _, y := range a[:i] {
			d := x - y
			for j >= 0 && a[j] > d {
				j--
			}
 
			if j >= 0 {
				ans = min(ans, d-a[j])
			}
			if j+1 < n {
				ans = min(ans, a[j+1]-d)
			}
		}
	}
 
	fmt.Fprintln(out, ans)
}