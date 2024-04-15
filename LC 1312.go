func minInsertions(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][i] = 1
    }

    for d := 2; d <= n; d++ {
        for i := 0; i <= n - d; i++ {
…        }
    }

    return n - dp[0][n-1]
}