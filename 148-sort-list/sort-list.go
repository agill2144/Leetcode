/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
    merge sort
    - dont have size of LL like we do with arrays, so we have to find mid
    - to get mid; we can use slow and fast ptrs
    - now we have to physically split the list into 2 halves
    - left and right half
    - split at mid
    - head -> mid = left
    - mid.Next -> end = right half
    - now recurse with 2 halves ( further split each half into more chunks )
    - at one point we will have a list of size 1 or 0
    - and that list we cannot further split, therefore its sorted on its own, return that 1 size list as is.
    - this way at one point, deep in our recursion tree
    - we will have 2 halves left and right ( visualize this at the bottom of the recursive tree )
    - now to merge 2 sorted lists, we can use "merge 2 sorted list" code 
        - using sentinel header
        - and appending smaller val first
    - then we can return the merged list back to parent recursive call
    - which will then repeat the same thing 

    time = o(nlogn); merge sort
        +  o(n)    ; finding mid
        +  o(n)    ; merged 2 sorted lists
    -----------------------------------------
        o(nlogn) + o(2n) = o(nlogn)
    
    space = o(logn) for the recursive stack
*/
func sortList(head *ListNode) *ListNode {
    // 4. when we only have 1 or 0 node, we dont have anything to split
    if head == nil || head.Next == nil {return head}

    // 1. get mid
    mid := mid(head)

    // 2. split into 2 halves
    left := head
    right := mid.Next
    mid.Next = nil

    // 3. recurse and let it split further
    left = sortList(left)
    right = sortList(right)

    // 5. merge 2 sorted list; merge left and right
    return merge2SortedLists(left, right)
}

// two ptrs
// time = o(n)
// space = o(1)
func merge2SortedLists(list1, list2 *ListNode) *ListNode {
    l1 := list1
    l2 := list2
    out := &ListNode{Val: 0}
    tail := out
    for l1 != nil && l2 != nil {
        
        l1Val := l1.Val
        l2Val := l2.Val
        
        if l1Val < l2Val {
            l1 = l1.Next
            tail.Next = &ListNode{Val: l1Val}
        } else {
            l2 = l2.Next
            tail.Next = &ListNode{Val: l2Val}
        }
        tail = tail.Next
    }
    
    for l1 != nil {
        tail.Next = &ListNode{Val: l1.Val}
        l1 = l1.Next
        tail = tail.Next
    }
    for l2 != nil {
        tail.Next = &ListNode{Val: l2.Val}
        l2 = l2.Next
        tail = tail.Next
    }
    return out.Next
}


/*
    - get mid using tortoise and hare method
    - slow ptr = tortoise
    - fast ptr = hare
    - slow moves 1x speed
    - fast moves 2x speed
    - if fast has reached the finish line
    - and slow is moving half the speed of fast
    - then slow ptr is half way across
    - meaning slow ptr is at our mid node
    time = o(n)
    space = o(1)
*/
func mid(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {return head}
    var prev *ListNode
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    return prev
}