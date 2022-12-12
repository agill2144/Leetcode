/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
    level order traversal.
    
    approach: BFS
    - maintain a level var ( start at 0)
    - the level var maps to idx in the result array
    - if the len of result == level, it means this levelIdx does not exist yet, so append
    - if not, it means that this level idx already exists, update value at levelIdx if currentVal > result[levelIdx]
    - level++ ( once a level is done ) so that the our level idx is pointing to next idx.
    time : o(n)
    space: o(maxWidthOfTree) or o(n/2)
    

    approach: DFS
    - level order is also possible using DFS, by maintaining a level value at each recursion call.
    - Same logic as the BFS
    time: o(n) - we visit all the nodes in the tree
    space: o(h) for avg case and o(n) for a skewed tree

*/


// DFS
func largestValues(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    result := []int{}
    var dfs func(level int, a *TreeNode)
    dfs = func(level int, a *TreeNode) {
        if a == nil {
            return
        }
        if len(result) == level {
            result = append(result, a.Val)
        } else if a.Val > result[level] {
            result[level] = a.Val
        }
        dfs(level+1, a.Left)
        dfs(level+1, a.Right)
    }
    dfs(0, root)
    return result
}


// BFS
// func largestValues(root *TreeNode) []int {
//     if root == nil {
//         return nil
//     }
//     result := []int{}
//     level := 0   
//     q := []*TreeNode{root}
//     for len(q) != 0 {
       
//         qSize := len(q)
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             if len(result) == level {
//                 result = append(result, dq.Val)
//             } else if dq.Val > result[level]{
//                 result[level] = dq.Val
//             }
            
//             if dq.Left != nil { q = append(q, dq.Left)}
//             if dq.Right != nil { q = append(q, dq.Right)}
            
//             qSize--
//         }
//         level++
    
//     }
//     return result
//  }