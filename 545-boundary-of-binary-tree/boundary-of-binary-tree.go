/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
    0 = root
    1 = left
    2 = right
    3 = middle
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



func leftID(r *TreeNode, currID int) int {
    if currID == 0 || currID == 1 {
        // root's left will always be a left boundary node
        // or 
        // an current left boundary node's left will always be a left node!
        return 1
    } else if currID == 2 {
        if r.Right == nil {
            return 2
        }
    }
    return 3
}

func rightID(r *TreeNode, currID int) int {
    if currID == 0 || currID == 2{
        // root's right will always be a right boundary node
        // or 
        // an current right boundary node's right will always be a right node!
        return 2
    } else if currID == 1 {
        if r.Left == nil {
            return 1
        }
    }
    return 3
} 