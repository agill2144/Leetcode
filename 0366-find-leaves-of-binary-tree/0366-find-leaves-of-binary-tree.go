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
func findLeaves(root *TreeNode) [][]int {
    result := [][]int{}
    var dfs func(r *TreeNode) int 
    dfs = func(r *TreeNode) int {
        // base
        // height of a nil node is 0
        if r == nil {return 0}
        
        // logic
        left := dfs(r.Left)
        right := dfs(r.Right)
        // we use maxH to use as an idx in result array
        maxH := max(left, right)
        if len(result) == maxH {
            result = append(result, []int{})
        }
        result[maxH] = append(result[maxH], r.Val)
        
        // max height of curr node is; max(left, right)+1
        return max(left,right)+1
    }
    dfs(root)
    return result
}

func max(x, y int) int {
    if x > y {return x}
    return y
}



// brute
// simply simulate/follow the example
// collect curr leaves
// return back whether a node should be deleted by parent or not
// a node can only be deleted by a parent if that node is a leaf node
// a parent will itself whether its a leaf or not FIRST
// then parent will check whether left / right child needs to be deleted ( delete them if yes , child return values will tell us that )
// if left child is a leaf, save its value, delete left
// if right child is a leaf, save its value, delete right
// return leaves collected and whether before deleting childs was this node a leaf node

// time: o(n^2)
// space: o(n) 
// func findLeaves(root *TreeNode) [][]int {
//     out := [][]int{}
    
//     var dfs func(r *TreeNode) ([]int, bool)
//     dfs = func(r *TreeNode) ([]int, bool) {
//         // base
//         if r == nil {return nil, false}
        
//         // logic
                
        
//         leftLeaves, leftIsLeaf := dfs(r.Left)
//         rightLeaves, rightIsLeaf := dfs(r.Right)
//         currIsLeaf := r.Left == nil && r.Right == nil
        
//         currLeaves := []int{}
//         if leftIsLeaf {currLeaves = append(currLeaves, r.Left.Val); r.Left = nil}
//         if rightIsLeaf {currLeaves = append(currLeaves, r.Right.Val); r.Right = nil}
//         currLeaves = append(currLeaves, leftLeaves...)
//         currLeaves = append(currLeaves, rightLeaves...)
        
//         return currLeaves, currIsLeaf 
//     }
    
//     for root.Left != nil || root.Right != nil {
//         leaves, _ := dfs(root)
//         out = append(out, leaves)
//     }
//     out = append(out, []int{root.Val})
    
//     return out
// }






