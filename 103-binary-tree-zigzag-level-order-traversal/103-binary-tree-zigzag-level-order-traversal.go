/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverse(a []int) {
    var l, r int = 0, len(a) - 1
    for l < r  {
        t := a[l]
        a[l] = a[r]
        a[r] = t
        l++
        r--
    }
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    q := make([]*TreeNode, 1)
    q[0] = root
    res := make([][]int, 0)
    odd := true
    for len(q) > 0 {
        size := len(q)
        layer := make([]int, 0)
        for i := 0; i < size; i++ {
            r := q[0]
            q = q[1 : ]
            if r.Left != nil {
                q = append(q, r.Left)
            }
            if r.Right != nil {
                q = append(q, r.Right)
            }
            layer = append(layer, r.Val)
        }
        if !odd {
            reverse(layer)
        } 
        odd = !odd
        res = append(res, layer)
    }
    return res
}