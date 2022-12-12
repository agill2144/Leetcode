/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    max depth as in max height of a branch.
    Left and right are going to return $x and $y as int
    and we have to return the max between both and add 1 to consider current node in max height
    
    We would like to go all the way down left and make left return its height to left leaf
    and then go right and return its right height
    
    Then once we are back the left leaf node we can compare the two and add + 1

*/
func maxDepth(root *TreeNode) int {
    return postOrderDfs(root)
}

func postOrderDfs(root *TreeNode) int {
    // base 
    if root == nil {
        return 0
    }
    
    // logic
    left := postOrderDfs(root.Left)    
    right := postOrderDfs(root.Right)
    return int(math.Max(float64(left), float64(right)))+1
}