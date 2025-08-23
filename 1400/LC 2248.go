func intersection(nums [][]int) []int {
    m := make(map[int]bool)
    for i, v := range nums {
        tmp := make(map[int]bool)
        for _, x := range v {
            if m[x] || i == 0 {
                tmp[x] = true
            }
        }
        m = tmp
    }

    ans := make([]int, 0, len(m))
    for k := range m {
        ans = append(ans, k)
    }
    sort.Ints(ans)

    return ans
}