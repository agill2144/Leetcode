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
    nums := []int{}    
    var dfs func(r *TreeNode)
    dfs = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        dfs(r.Left)
        nums = append(nums, r.Val)
        dfs(r.Right)
    }
    dfs(root)
    sort.Ints(nums)
    ptr := 0
    var writeBack func(r *TreeNode)
    writeBack = func(r *TreeNode) {
        // base
        if r == nil {return}

        // logic
        writeBack(r.Left)
        r.Val = nums[ptr]
        ptr++
        writeBack(r.Right)
    }
    writeBack(root)
}