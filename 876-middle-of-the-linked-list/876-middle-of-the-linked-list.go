/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    a := head
    b := head.Next
    for b != nil {
        a = a.Next
        b = b.Next
        if b != nil {
            b = b.Next
        }
    }
    return a
}