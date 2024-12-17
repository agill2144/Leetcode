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
    reservoir := this.head // Reservoir initially holds the first node
    count := 1        // Start with the first node (so count is 1)
    
    // Traverse the linked list
    current := this.head.Next
    for current != nil {
        count++
        
        // Generate a random number between 0 and count-1
        randNum := rand.Intn(count)
        
        // With probability 1/count, replace the reservoir
        if randNum == 0 {
            reservoir = current
        }
        
        current = current.Next
    }
    
    return reservoir.Val // Return the randomly selected node's value
    
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */