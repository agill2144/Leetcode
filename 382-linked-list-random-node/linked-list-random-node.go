/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    head *ListNode
}


func Constructor(head *ListNode) Solution {
    return Solution{head}
}


func (this *Solution) GetRandom() int {
    res := -1
    count := 0
    curr := this.head
    for curr != nil {
        count++
        if rand.Intn(count) == 0 {
            res = curr.Val
        }
        curr = curr.Next
    }
    return res
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */