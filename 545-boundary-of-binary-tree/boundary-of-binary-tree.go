/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func boundaryOfBinaryTree(root *TreeNode) []int {
    left := []int{}
    leaves := []int{}
    right := []int{}
    var dfs func(r *TreeNode, currRole int) 
    dfs = func(r *TreeNode, currRole int) {
        // base
        if r == nil {return}

        // logic
        if r.Left == nil && r.Right == nil {
            leaves = append(leaves, r.Val)
        } else if currRole == 0 || currRole == 1 {
            left = append(left, r.Val)
        } else if currRole == 2 {
            right = append(right, r.Val)
        }
        dfs(r.Left, leftChild(r, currRole))
        dfs(r.Right, rightChild(r, currRole))
    }
    dfs(root, 0)
    left = append(left, leaves...)
    for i := len(right)-1; i >= 0; i-- {
        left = append(left, right[i])
    }
    return left
}

func leftChild(r *TreeNode, currRole int) int {
    if currRole == 0 || currRole == 1 {
        return 1
    } else if currRole == 2 {
        if r.Right == nil {
            return 2
        }
    }
    return 3
}

func rightChild(r *TreeNode, currRole int) int {
    if currRole == 0 || currRole == 2 {
        return 2
    } else if currRole == 1 {
        if r.Left == nil {
            return 1
        }
    }
    return 3
}