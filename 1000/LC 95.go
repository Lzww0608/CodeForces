/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func clone(root *TreeNode, offset int) *TreeNode {
    if root == nil {
        return nil
    }

    cur := &TreeNode{root.Val + offset, clone(root.Left, offset), clone(root.Right, offset)}
    return cur
}

func generateTrees(n int) []*TreeNode {
    ans := make([][]*TreeNode, n + 1)
    ans[0] = append(ans[0], nil)
    for d := 1; d <= n; d++ {
        for root := 1; root <= d; root++ {
            left := root - 1
            right := d - root
            for _, p := range ans[left] {
                for _, q := range ans[right] {
                    cur := &TreeNode{root, clone(p, 0), clone(q, root)}
                    ans[d] = append(ans[d], cur)
                }
            }
        }
    }

    return ans[n]
}