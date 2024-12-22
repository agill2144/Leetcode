/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    items []int
}


func Constructor(head *ListNode) Solution {
    items := []int{}
    for head != nil {
        items = append(items, head.Val)
        head = head.Next
    }
    return Solution{items}    
}


func (this *Solution) GetRandom() int {
    n := len(this.items)
    r := rand.Intn(n)
    return this.items[r]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */