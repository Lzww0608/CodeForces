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

	var n, x int
	fmt.Fscan(in, &n)
	cnt := make(map[int][]int)
	for i := 0; i < n<<1; i++ {
		fmt.Fscan(in, &x)
		cnt[x] = append(cnt[x], i+1)
	}

	for _, v := range cnt {
		if len(v)&1 == 1 {
			fmt.Fprintln(out, -1)
			return
		}
	}

	for _, v := range cnt {
		for i := 0; i < len(v); i += 2 {
			fmt.Fprintln(out, v[i], v[i+1])
		}
	}
	return
}
