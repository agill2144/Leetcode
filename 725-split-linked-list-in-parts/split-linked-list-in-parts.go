/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
    approach: linear scan + ptrs
    - we need a k size output list
    - each ith element is supposed to contains a x size LL
    - each x size depends on the full size of the LL
    - "length of each part should be as equal as possible"
    - meaning if n = 10; and k = 2 
    - then, we have to create a 2 size output list with 5 nodes in each (n/k)
    - what if n was not perfectly divisible by k ?
    - n = 11 and k = 3; 11/3 = 3; and we have remainder = 2
        - "parts occurring earlier should always have a size greater"
    - so the first 2 lists in output array, each of them will get 1 extra node
    - meaning, even when we have remainder, we need to divide them equally starting from the start of the output array

    - NOTE: the remainder will NEVER exceed k-1
    - n = 40 ; k = 3; 40/3 = 13 and remainder = 1 (40%3)
    - n = 41 ; k = 3; 41/3 = 13 and remainder = 2 (41%3)
    - n = 42 ; k = 3; 42/3 = 14 and remainder = 0 (42%3)
    - therefore we wont have remainder higher than size of output array
        and we wont have to wrap around the output array distributing 
        remaining nodes. because the remaining nodes never exceed
        k-1, which is less than the size of our output array
    
    - take a pass to calculate size of LL ( n )
    - then figure out the size of each split (n/k)
    - then figure out the remaining nodes we need to add starting from beginning (n%k)
    - then take a curr ptr and loop until end of linkedlist
    - set

*/
func splitListToParts(head *ListNode, k int) []*ListNode {
    out := make([]*ListNode, k)
    if head == nil {return out}
    n := 0
    curr := head
    for curr != nil {
        n++
        curr = curr.Next
    }
    grpSize := n / k
    rem := n % k
    curr = head
    outPtr := 0
    size := 0
    tail := curr
    start := curr
    for curr != nil {
        if size == grpSize {
            // take the extra node if we need to
            if rem != 0 && curr != nil {
                tail = curr
                curr = curr.Next
                rem--
            }
            tail.Next = nil
            out[outPtr] = start
            size = 0
            start = curr
            tail = curr
            outPtr++
        } else {
            size++
            tail = curr
            curr = curr.Next
        }
    }
    if size != 0{
        out[outPtr] = start
    }
    return out
}