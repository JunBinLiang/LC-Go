/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int
func dfs(root *TreeNode) []int {
    if root == nil {
        return []int{0, 0}
    }
    l := dfs(root.Left)
    r := dfs(root.Right)
    sum := l[0] + r[0] + root.Val
    cnt := l[1] + r[1] + 1
    if sum / cnt == root.Val {
        res++
    }
    return []int{sum, cnt}
}

func averageOfSubtree(root *TreeNode) int {
    res = 0
    dfs(root)
    return res
}