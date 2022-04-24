import (
    "fmt"
    "container/heap"
)


type IntHeap []int 
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}    
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type KthLargest struct {
    k  int 
    h  *IntHeap
}


func Constructor(k int, nums []int) KthLargest {
    h := &IntHeap{}
    
    for i := 0; i < len(nums); i++ {
        heap.Push(h, nums[i])
    }
    
    for h.Len() > k {
        heap.Pop(h)
    }
    
    return KthLargest{k, h}
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.h, val)
    if this.h.Len() > this.k {
        heap.Pop(this.h)
    }
    return (*this.h)[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */