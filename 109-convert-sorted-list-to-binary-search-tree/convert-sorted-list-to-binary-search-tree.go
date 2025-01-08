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
    n := size(head)
    var dfs func(left, right int) *TreeNode
    dfs = func(left, right int) *TreeNode {
        // base
        if left > right {return nil}

        // logic
        mid := left + (right-left)/2
        leftChild := dfs(left, mid-1)
        if head != nil {
            root := &TreeNode{Val: head.Val}
            head = head.Next
            root.Left = leftChild
            root.Right = dfs(mid+1, right)
            return root
        }
        return nil
    }
    return dfs(0, n-1)
}

func size(head *ListNode) int {
    n := 0
    curr := head
    for curr != nil {
        n++
        curr = curr.Next
    }
    return n
}