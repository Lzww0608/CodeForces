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
 
	var t, n int
next:
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		m := map[int]bool{}
		for _, x := range a {
			if m[x] {
				fmt.Fprintln(out, "YES")
				continue next
			}
			m[x] = true
		}
		fmt.Fprintln(out, "NO")
	}
}