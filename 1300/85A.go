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

	var n int
	fmt.Fscan(in, &n)
	if n&1 == 1 {
		for i := 0; i < 4; i++ {
			if i == 0 || i == 1 {
				for j := 0; j < n-1; j += 2 {
					if (i&1+(j/2)&1)&1 == 0 {
						fmt.Fprint(out, "aa")
					} else {
						fmt.Fprint(out, "bb")
					}
				}
				fmt.Fprintln(out, "e")
			} else {
				fmt.Fprint(out, "f")
				for j := 1; j < n; j += 2 {
					if (i&1+((j-1)/2)&1)&1 == 0 {
						fmt.Fprint(out, "cc")
					} else {
						fmt.Fprint(out, "dd")
					}
				}
				fmt.Fprintln(out)
			}
		}
		return
	}
	for i := 0; i < 4; i++ {
		if i == 0 || i == 3 {
			for j := 0; j < n; j += 2 {
				if (i&1+(j/2)&1)&1 == 0 {
					fmt.Fprint(out, "aa")
				} else {
					fmt.Fprint(out, "bb")
				}
			}
			fmt.Fprintln(out)
		} else {
			fmt.Fprint(out, "e")
			for j := 1; j < n-1; j += 2 {
				if (i&1+((j-1)/2)&1)&1 == 0 {
					fmt.Fprint(out, "cc")
				} else {
					fmt.Fprint(out, "dd")
				}
			}
			fmt.Fprintln(out, "f")
		}
	}
}
