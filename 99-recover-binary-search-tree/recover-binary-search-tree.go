/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    approach: brute force
    - traverse inorder
    - write to a list
    - sort the list
    - use a ptr on list and traverse the bst in inorder
    - while processing a root node, write the value from sortedList[ptr] to current root node

    tc = o(n) + o(nlogn) + o(n)
    sc = o(n)
*/
func recoverTree(root *TreeNode)  {
    if root == nil {return}
    var (
        fb *TreeNode
        sb *TreeNode
    )
    var prev *TreeNode
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Left)
        if prev != nil {
            if prev.Val >= r.Val {
                if fb == nil {fb = prev}
                sb = r
            }
        }
        prev = r
        dfs(r.Right)
    }
    dfs(root)
    fb.Val, sb.Val = sb.Val, fb.Val
    
}