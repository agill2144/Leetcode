/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: brute force with extra space
    - traverse the LL and put the node values into a list
    - then build a height balanced bst using the sorted list
    - the root node will the mid element; given 2 idxs left and right 
    - the left of this root will be elements from left -> mid-1
    - the right of this root will be elements from mid+1 -> right

    tc = o(2n)
    sc = o(n)

    approach: no extra list space
    - we can get mid of linkedlist using slow and fast ptrs
    - and then split the list into 2 halves; for creating left subtree and for creating right subtree

*/
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {return nil}
    var dfs func(left, right *ListNode) *TreeNode
    dfs = func(left, right *ListNode) *TreeNode {
        // base
        if left == right {return nil}

        // logic
        slow := left
        fast := left
        for fast != right && fast.Next != right {
            slow = slow.Next
            fast = fast.Next.Next
        }
        // slow is at mid point
        root := &TreeNode{Val: slow.Val}
        root.Left = dfs(left, slow)
        root.Right = dfs(slow.Next, right)
        return root
    }
    return dfs(head, nil)
    
}