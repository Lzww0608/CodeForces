const MOD int = 1_000_000_007
const N int = 505
var f [N]int
func init() {
    f[0] = 1
    for i := 1; i < N; i++ {
        for j := 0; i - 1 - j >= 0; j++ {
            f[i] += f[j] * f[i-1-j]
            f[i] %= MOD
        }
    }
}
func numberOfWays(numPeople int) int {
    return f[numPeople >> 1]
}