/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// /*
//     pattern: bottom up, tree height
//     approach: use height as idx in result array for each node
//     time: o(n)
//     space: o(n) worst case we have a skewed tree
// */
// func findLeaves(root *TreeNode) [][]int {
//     result := [][]int{}
//     var dfs func(r *TreeNode) int
//     dfs = func(r *TreeNode) int {
//         // base
//         if r == nil {return 0}
        
//         // logic
//         left := dfs(r.Left)
//         right := dfs(r.Right)
//         maxHeight := max(left, right)
//         if len(result) == maxHeight {
//             result = append(result, []int{})
//         }
//         result[maxHeight] = append(result[maxHeight],r.Val)
//         return maxHeight+1
//     }
//     dfs(root)
//     return result
// }

// func max(x, y int) int {
//     if x > y {return x}
//     return y
// }



func findLeaves(root *TreeNode) [][]int {
    out := [][]int{}
    
    var dfs func(r *TreeNode) ([]int, bool)
    dfs = func(r *TreeNode) ([]int, bool) {
        // base
        if r == nil {return nil, false}
        
        // logic
        path := []int{}
        if r.Left == nil && r.Right == nil {
            path = append(path, r.Val)
            return path, true
        }        
        
        leftLeaves, leftIsLeaf := dfs(r.Left)
        rightLeaves, rightIsLeaf := dfs(r.Right)
        currIsLeaf := r.Left == nil && r.Right == nil
        
        path = append(path, leftLeaves...)
        path = append(path, rightLeaves...)
        if leftIsLeaf {r.Left = nil}
        if rightIsLeaf {r.Right = nil}
        
        return path, currIsLeaf 
    }
    
    for root.Left != nil || root.Right != nil {
        leaves, _ := dfs(root)
        out = append(out, leaves)
    }
    out = append(out, []int{root.Val})
    
    return out
}






