package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 26

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	fmt.Fscan(in, &s)
	cnt := [N]int{}
	for i := range s {
		cnt[int(s[i]-'a')]++
	}

	odd := 0
	for _, x := range cnt {
		if x&1 == 1 {
			odd++
		}
	}

	if odd&1 == 1 || odd == 0 {
		fmt.Fprintln(out, "First")
	} else {
		fmt.Fprintln(out, "Second")
	}
}
