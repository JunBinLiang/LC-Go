/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int = 0
func dfs(root *TreeNode, mx int) {
    if root == nil {
        return
    }
    if root.Val >= mx {
        res++
        mx = root.Val
    }
    dfs(root.Left, mx)
    dfs(root.Right, mx)
}

func goodNodes(root *TreeNode) int {
    res = 0
    dfs(root, -100000)
    return res
}