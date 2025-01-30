/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
    id = 0 = root
    id = 1 = left
    id = 2 = right
    id = 2 = middle
*/
func boundaryOfBinaryTree(root *TreeNode) []int {
    if root == nil {return nil}
    left := []int{}
    leaves := []int{}
    right := []int{}
    var dfs func(r *TreeNode, id int)
    dfs = func(r *TreeNode, id int) {
        // base
        if r == nil {return}

        // logic
        if r.Left == nil && r.Right == nil {
            leaves = append(leaves, r.Val)
        } else if id == 0 || id == 1 {
            left = append(left, r.Val)
        } else if id == 2 {
            right = append(right, r.Val)
        }
        dfs(r.Left, leftID(r,id))
        dfs(r.Right, rightID(r,id))
    }
    dfs(root, 0)
    left = append(left, leaves...)
    for i := len(right)-1; i >= 0; i-- {
        left = append(left, right[i])
    }
    return left
}


func leftID(r *TreeNode, id int) int {
    if id == 0 || id == 1 {
        return 1
    } else if id == 2 {
        if r.Right == nil {
            return 2
        }
    }
    return 3
}

func rightID(r *TreeNode, id int) int {
    if id == 0 || id == 2 {
        return 2
    } else if id == 1 {
        if r.Left == nil {
            return 1
        }
    }
    return 3
}