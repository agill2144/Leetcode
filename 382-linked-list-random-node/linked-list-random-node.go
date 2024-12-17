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
    curr := this.head
    count := 0
    for curr != nil {
        count++
        randNum := rand.Intn(count)        
        // With probability 1/count, replace the reservoir
        if randNum == 0 {
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