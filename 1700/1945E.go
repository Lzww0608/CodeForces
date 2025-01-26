package main
 
import (
	"bufio"
	"fmt"
	"os"
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
	pos := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		if a[i] == k {
			pos = i
		}
	}
 
	l, r := 0, n
	for l+1 < r {
		mid := l + ((r - l) >> 1)
		if a[mid] <= k {
			l = mid
		} else {
			r = mid
		}
 
	}
 
	fmt.Fprintln(out, 1)
	fmt.Fprintln(out, pos+1, l+1)
}