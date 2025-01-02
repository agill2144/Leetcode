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
	var dfs func(h *ListNode) *ListNode
	dfs = func(h *ListNode) *ListNode {
		// base
		// if this is the only node, return it as is!
		if h == nil || h.Next == nil {
			return h
		}

		// logic
		var prev *ListNode
		slow := h
		fast := h
		for fast != nil && fast.Next != nil {
			prev = slow
			slow = slow.Next
			fast = fast.Next.Next
		}
		// create 2 halves by splitting them
		prev.Next = nil
		// recurse on the 2 halves
		left := dfs(h)
		right := dfs(slow)
		// merge 2 halves using merge2SortedLists(l1, l2)
		// and return the merged result back to parent
		return merge2SortedLists(left, right)
	}
	return dfs(head)
}

func merge2SortedLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	tail := dummy
	for l1 != nil || l2 != nil {
		l1Val := math.MaxInt64
		l2Val := math.MaxInt64
		if l1 != nil {
			l1Val = l1.Val
		}
		if l2 != nil {
			l2Val = l2.Val
		}
		if l1Val <= l2Val {
			tail.Next = l1
			tail = tail.Next
			l1 = l1.Next
		} else {
			tail.Next = l2
			tail = tail.Next
			l2 = l2.Next
		}
	}
	return dummy.Next
}
