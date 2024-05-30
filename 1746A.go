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
 
	var t, n, k int
	for fmt.Fscan(in, &t); t > 0; t-- {
		fmt.Fscan(in, &n, &k)
		var x int
		var f bool = false
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &x)
			if x == 1 {
				f = true
			}
		}
		if f {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
 
	}
}