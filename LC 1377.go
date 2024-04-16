func frogPosition(n int, edges [][]int, t int, target int) float64 {
    g := make([][]int, n + 1)
    for _, e := range edges {
        a, b := e[0], e[1]
        g[a] = append(g[a], b)
        g[b] = append(g[b], a)
    }
    g[1] = append(g[1], 0)
    vis := make([]bool, n + 1)
    //var prob float64 = 1.0
    m := map[int]float64{}

    var dfs func(int, int, float64) 
    dfs = func(x, time int, prob float64) {
        if vis[x] == true || time > t {
            return 
        }
        vis[x] = true
        d := len(g[x]) - 1
        if d == 0 || time == t {
            m[x] = prob
            return 
        }

        for _, t := range g[x] {
            dfs(t, time + 1, prob / float64(d))
        }
    }
    dfs(1, 0, 1.0)

    return m[target]
}