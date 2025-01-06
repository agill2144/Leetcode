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
func isSubPath(head *ListNode, root *TreeNode) bool {
    if head == nil && root == nil {return true}
    if head == nil || root == nil {return false}
    var dfs func(r *TreeNode, curr *ListNode) bool
    dfs = func(r *TreeNode, curr *ListNode) bool {
        // base
        if curr == nil {return true}
        if r == nil {return false}

        // logic
        if curr.Val == r.Val  {
            return dfs(r.Left, curr.Next) || dfs(r.Right, curr.Next)
        }
        return false
    }
    return dfs(root, head) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}