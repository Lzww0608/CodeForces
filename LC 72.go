func minDistance(s string, t string) int {
    m, n := len(s), len(t)
    f := make([][]int, m + 1)
    for i := range f {
        f[i] = make([]int, n + 1)
    }
    for j := 0; j <= n; j++ {
        f[0][j] = j
    }
    
    for i := range s {
        f[i+1][0] = i + 1
        for j := range t {
            if s[i] == t[j] {
                f[i+1][j+1] = f[i][j]
            } else{
                f[i+1][j+1] = min(f[i][j], f[i+1][j], f[i][j+1]) + 1 
            }
        }
    }

    return f[m][n]
}