package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

const N int = 26

var id []int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	id = make([]int, N)
	for i := range id {
		id[i] = i
	}

	var t int
	for fmt.Fscan(in, &t); t > 0; t-- {
		solve(in, out)
	}
}

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	var s []byte
	fmt.Fscan(in, &n, &s)
	cnt := make([]int, N)

	for i := range s {
		cnt[int(s[i]-'a')]++
	}

	sort.Slice(id, func(i, j int) bool {
		return cnt[id[i]] > cnt[id[j]]
	})

	maxSave, k := 0, 0
	for i := 1; i <= N; i++ {
		if n%i != 0 {
			continue
		}
		save := 0
		for _, j := range id[:i] {
			save += min(cnt[j], n/i)
		}
		if maxSave < save {
			maxSave = save
			k = i
		}
	}

	need := []byte{}
	for _, j := range id[:k] {
		if cnt[j] >= n/k {
			cnt[j] = n / k
		} else {
			need = append(need, bytes.Repeat([]byte{byte(j) + 'a'}, n/k-cnt[j])...)
		}
	}

	for _, j := range id[k:] {
		cnt[j] = 0
	}

	for i := 0; i < n; i++ {
		x := int(s[i] - 'a')
		if cnt[x] > 0 {
			cnt[x]--
		} else {
			s[i] = need[0]
			need = need[1:]
		}
	}

	fmt.Fprintln(out, n-maxSave)
	fmt.Fprintln(out, string(s))
}
