/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge(h1 *ListNode, h2 *ListNode) *ListNode {
    dummy := &ListNode{}

    cur := dummy
    for h1 != nil && h2 != nil {
        if h1.Val < h2.Val {
            cur.Next = h1
            h1 = h1.Next
        } else {
            cur.Next = h2
            h2 = h2.Next
        }
        cur = cur.Next
    }
    
    for h1 != nil {
        cur.Next = h1
        h1 = h1.Next
        cur = cur.Next
    }
    
    for h2 != nil {
        cur.Next = h2
        h2 = h2.Next
        cur = cur.Next
    }
    
    return dummy.Next
}

func divide(lists []*ListNode, l int, r int) *ListNode {
    if l == r {
        return lists[l];
    }
    
    mid := l + (r - l) / 2
    a := divide(lists, l, mid)
    b := divide(lists, mid + 1, r)
    return merge(a, b)
}
 
func mergeKLists(lists []*ListNode) *ListNode {
    if lists == nil || len(lists) == 0 {
        return nil
    }
    return divide(lists, 0, len(lists) - 1)
}