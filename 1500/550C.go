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

	var s string
	fmt.Fscan(in, &s)
	n := len(s)
	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		if x%8 == 0 {
			fmt.Fprintln(out, "YES")
			fmt.Fprintln(out, x)
			return
		}
		for j := i + 1; j < n; j++ {
			y := int(s[j] - '0')
			if ((x&15)*10+y)%8 == 0 {
				fmt.Fprintln(out, "YES")
				fmt.Fprintln(out, x*10+y)
				return
			}
			for k := j + 1; k < n; k++ {
				z := int(s[k] - '0')
				if ((x&15)*100+(y&15)*10+z)%8 == 0 {
					fmt.Fprintln(out, "YES")
					fmt.Fprintln(out, x*100+y*10+z)
					return
				}
			}
		}
	}

	fmt.Fprintln(out, "NO")
}
