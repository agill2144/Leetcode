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
    var dfs func(curr *ListNode) *ListNode
	dfs = func(curr *ListNode) *ListNode {
		// base
		if curr == nil || curr.Next == nil {
			return curr
		}

		// logic
		slow := curr
		fast := curr
		var prev *ListNode
		for fast != nil && fast.Next != nil {
			prev = slow
			slow = slow.Next
			fast = fast.Next.Next
		}
		prev.Next = nil
		return mergeTwoLists(dfs(curr), dfs(slow))
	}
	return dfs(head)
}

func mergeTwoLists(x *ListNode, y *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	tail := dummy
	p1, p2 := x, y
	for p1 != nil || p2 != nil {
		p1Val := math.MaxInt64
		if p1 != nil { p1Val = p1.Val }
		p2Val := math.MaxInt64
		if p2 != nil { p2Val = p2.Val }
		if p1Val <= p2Val {
			tail.Next = p1
			tail = tail.Next
			p1 = p1.Next
		} else {
			tail.Next = p2
			tail = tail.Next
			p2 = p2.Next
		}
	}
	return dummy.Next
}

