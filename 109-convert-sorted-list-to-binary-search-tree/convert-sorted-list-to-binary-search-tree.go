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
        root := &TreeNode{Val: slow.Val}
        root.Left = dfs(left, slow)
        root.Right = dfs(slow.Next, right)
        return root
    }
    return dfs(head, nil)

}