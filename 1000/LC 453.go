func minMoves(nums []int) (ans int) {
    mn := slices.Min(nums)
    for _, x := range nums {
        ans += x - mn
    }

    return
}