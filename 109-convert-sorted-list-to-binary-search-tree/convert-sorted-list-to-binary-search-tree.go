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
    var dfs func(curr *ListNode) *TreeNode
    dfs = func(curr *ListNode) *TreeNode {
        // base
        if curr == nil {return nil}
        if curr.Next == nil {
            return &TreeNode{Val: curr.Val}
        }        

        // logic
        prev := curr
        slow := curr
        fast := curr
        for fast != nil && fast.Next != nil {
            prev = slow
            slow = slow.Next
            fast = fast.Next.Next
        }
        // slow ptr is at middle node
        // therefore mid is our root node
        root := &TreeNode{Val:slow.Val}
        // cut the list into 2 halves
        // left half is the left subtree (curr->prev)
        // right half is right subtree (slow.Next->right)
        // slow is already used! slow is the root node, therefore right half is after slow ptr
        prev.Next = nil
        root.Left = dfs(curr)
        root.Right = dfs(slow.Next)
        return root
    }
    return dfs(head)
}