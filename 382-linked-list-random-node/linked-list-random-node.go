/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


/*
    approach: brute force
    - save all node values in a list
    - use rand.Intn to select a random idx from list
        - each value has a equal probability of being selected    
    - return the value at that idx position
    tc = o(n)
    sc = o(n)


    Follow up:
    Could you solve this efficiently without using extra space?
    What if the linked list is extremely large and its length is unknown to you?
    - reservoir sampling!

    approach: reservoir sampling
    - reservoir = bucket of some size
        - we only care about 1 element
        - therefore size is 1
    - use counter to figure out whether current node 
        could replace our reservoir element
    - as counter goes up, probability of each element seen so far goes down
    - most importantly probability of current element being selected goes down
    - use rand.Intn(count) == 0 to check whether we can we can replace our reservoir elememt or not
        - 0 is a constant number, it will always be part of rand.Intn(count)
        - as rand.Intn(n) picks a number from [0, n)
        - therefore if random number is 0, with the lower probability counter, we can replace our res with current element

    tc = o(n)
    sc = o(1)

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
        if rand.Intn(count) == count-1 {
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