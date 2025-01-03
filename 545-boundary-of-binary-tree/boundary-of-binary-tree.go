/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    0 = root ( curr node is root )
    1 = left boundary node ( curr node is left boundary )
    2 = right boundary node ( curr node is right boundary )
    3 = middle node ( curr node is middle node )
*/
func boundaryOfBinaryTree(root *TreeNode) []int {
    left := []int{}
    right := []int{}
    leaves := []int{}
    var dfs func(r *TreeNode, state int)
    dfs = func(r *TreeNode, state int) {
        // base
        if r == nil {return}

        // logic
        if state == 0 {
            left = append(left, r.Val)
        } else if state == 1 {
            left = append(left, r.Val)
        } else if state == 2 {
            right = append(right, r.Val)
        } else if r.Left == nil && r.Right == nil {
            leaves = append(leaves, r.Val)
        }
        dfs(r.Left, leftState(r,state))
        dfs(r.Right, rightState(r,state))
    }
    dfs(root, 0)
    left = append(left, leaves...)
    for i := len(right)-1; i >= 0; i-- {
        left = append(left, right[i])
    }
    return left
}


func leftState(r *TreeNode, curr int) int {
    if curr == 0 || curr == 1 {
        return 1
    } else if curr == 2 {
        if r.Right == nil {return 2}
    }
    return 3
}

func rightState(r *TreeNode, curr int) int {
    if curr == 0 || curr == 2 {
        return 2
    } else if curr == 1 {
        if r.Left == nil {return 1}
    }
    return 3
}
