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
    curr := head
    items := []int{}
    for curr != nil {
        items = append(items, curr.Val)
        curr = curr.Next
    }
    return Solution{items}
}


func (this *Solution) GetRandom() int {
    r := rand.Intn(len(this.items))
    return this.items[r]    
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */