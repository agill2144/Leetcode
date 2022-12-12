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
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {return nil}
    
    var dfs func(left *ListNode, right *ListNode) *TreeNode
    dfs = func(left, right *ListNode) *TreeNode {
        // base
        if left == right {return nil}
        
        // logic
        // find mid given start = left and end = right
        slow := left
        fast := left
        for slow != nil && fast != right && fast.Next != right {
            slow = slow.Next
            fast = fast.Next.Next
        }
        // mid is at slow now
        root := &TreeNode{Val: slow.Val}
        // left of mid is the left subtree
        // meaning we need the next recursion call to find the mid between left - mid/slow
        root.Left = dfs(left, slow)
        root.Right = dfs(slow.Next, right)
        return root
    }
    return dfs(head, nil) 
}
