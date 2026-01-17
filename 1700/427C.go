package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://codeforces.com/problemset/problem/427/C
func CF427C_Kosaraju(_r io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	in := bufio.NewReader(_r)

	var n, m, v, w int
	Fscan(in, &n)

	// 读取花费
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	// 1. 建图：Kosaraju 需要原图 g 和反图 rg
	g := make([][]int, n)
	rg := make([][]int, n) // reverse graph
	for Fscan(in, &m); m > 0; m-- {
		Fscan(in, &v, &w)
		v-- // 转为 0-based
		w--
		g[v] = append(g[v], w)
		rg[w] = append(rg[w], v) // 反向边
	}

	// 2. 第一次 DFS：确定访问顺序 (对应你提供的 dfs1)
	visited := make([]bool, n)
	order := make([]int, 0, n) // 充当栈 s

	var dfs1 func(int)
	dfs1 = func(u int) {
		visited[u] = true
		for _, v := range g[u] {
			if !visited[v] {
				dfs1(v)
			}
		}
		// 递归结束后入栈，保证拓扑序靠前的后入栈
		order = append(order, u)
	}

	// 对所有未访问的点跑 dfs1 (处理非连通图)
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs1(i)
		}
	}

	// 3. 第二次 DFS：在反图上寻找 SCC 并统计答案 (对应你提供的 dfs2 + 统计逻辑)
	// 重置 visited 数组，或者使用一个新的数组来标记该点是否已被分配到某个SCC
	allocated := make([]bool, n)

	// 当前 SCC 的统计变量
	var sccMinVal int
	var sccMinCnt int

	var dfs2 func(int)
	dfs2 = func(u int) {
		allocated[u] = true

		// --- 题目特定逻辑：统计当前 SCC 的最小值 ---
		if a[u] < sccMinVal {
			sccMinVal = a[u]
			sccMinCnt = 1
		} else if a[u] == sccMinVal {
			sccMinCnt++
		}
		// ---------------------------------------

		for _, v := range rg[u] {
			if !allocated[v] {
				dfs2(v)
			}
		}
	}

	costSum := int64(0)
	ways := int64(1)

	// 按照 order 的逆序（栈顶开始）进行处理
	for i := len(order) - 1; i >= 0; i-- {
		u := order[i]
		if !allocated[u] {
			// 发现一个新的强连通分量
			// 初始化当前 SCC 的统计数据
			sccMinVal = int(1e9) + 7
			sccMinCnt = 0

			dfs2(u) // 在反图上跑，收集这个 SCC 的所有点信息

			// 累加结果
			costSum += int64(sccMinVal)
			ways = (ways * int64(sccMinCnt)) % mod
		}
	}

	Fprint(out, costSum, ways)
}

func main() { CF427C_Kosaraju(os.Stdin, os.Stdout) }
