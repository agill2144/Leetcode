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
    nodes := []int{}
    curr := head
    for curr != nil {nodes = append(nodes, curr.Val); curr = curr.Next}
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}
        // logic
        mid := left + (right-left)/2
        root := &TreeNode{Val: nodes[mid]}
        root.Left = dfs(left, mid-1)
        root.Right = dfs(mid+1, right)
        return root
    }
    return dfs(0, len(nodes)-1)
    
}