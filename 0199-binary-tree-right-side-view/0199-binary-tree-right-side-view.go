/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
    approach 1: level order using BFS
    This was the very first thing that came to my mind.
    Before processing each level ( in a queue ), we will save the last element from it
    and then process the nodes in queue like any other BFS

    time: o(n) - where n is the number of nodes in the tree. We visit and see every single node best/worst case ( we do not exit early anywhere )
    space: o(math.Max(widthOfTree)) or o(w) -- if its a skewed tree it becomes o(n)

    ------------------------------------------------

    approach 2: level order using DFS
    We can do level order traversal using DFS too by maintaining a level at each node in recursion stack.
    But how do we only save the right side of the tree?
    We will have a result array and each idx in this result array will represent a level in the tree
    level == idx and if len(res) == level this means the max idx avail right now is level-1 and if level == len(res) and level is our idx
    then that means we do not have level as an idx avail ( its less then level ) - so append in this case.
    
    2 ways:
    1. We can traverse left first and go reckless override each new element at the same level. 
        - if right goes after left, the last overrides will be the right side view
    2. We can traverse right first and only add to result array at the level IF the size of result array == level (i.e this is the first time we have seen this level)
        - gather all the right side values for each level
        - if this level does not exist in our result array
            - that means this is the first time we have seen this level, so append it with current level value
        - if this level does exist, it means we already have the farthest right side value for this level
    
    time: o(n) - where n is the number of nodes in the tree. We visit and see every single node best/worst case ( we do not exit early anywhere )
    space: o(h) - where h is the max height of tree and it will be o(n) space in a skewed tree
   
   
*/

// time: o(n) - we visit all the nodes
// space: o(w) - at worse our queue will have maxWidth of the tree saved in queue
// the maxWidth is usually the leaf level, or n/2 nodes where n is total number of nodes in tree
// level order using BFS
// func rightSideView(root *TreeNode) []int {
//     if root == nil {
//         return nil
//     }
//     result := []int{}
//     q := []*TreeNode{root}
    
//     for len(q) != 0 {
//         qSize := len(q)
//         result = append(result, q[qSize-1].Val)
//         for qSize != 0 {
//             dq := q[0]
//             q = q[1:]
//             if dq.Left != nil {q = append(q, dq.Left)}
//             if dq.Right != nil {q = append(q, dq.Right)}
//             qSize--
//         }
//     }
//     return result
// }


// time: o(n) - we visit all the nodes
// space: o(h) - at worse our recursion stack will have max height of the tree saved
// level order using DFS ( i.e maintain a level local state at each node )
// this exists to scope down global to only this class/struct or else global pollution happens
// func rightSideView(root *TreeNode) []int {
//     // right side first
//     var result []int
//     var dfs func(a *TreeNode, level int)
//     dfs = func(a *TreeNode, level int){
//         // base
//         if a == nil {return}
        
//         // logic
//         if len(result) == level {
//             // this means this level idx does not exist 
//             // and since we traverse right side first, save this item
//             result = append(result, a.Val)
//         }
//         dfs(a.Right, level+1)
//         dfs(a.Left, level+1)
//     }
//     dfs(root, 0)
//     return result
// }


// time: o(n) - we visit all the nodes
// space: o(h) - at worse our recursion stack will have max height of the tree saved
func rightSideView(root *TreeNode) []int {
    var result []int
    var dfs func(r *TreeNode, level int)
    dfs = func(r *TreeNode, level int) {
        // base
        if r == nil {return}
        
        // logic
        if len(result) == level {
            result = append(result, r.Val)
        } else {
            result[level] = r.Val
        }
        dfs(r.Left, level+1)
        dfs(r.Right, level+1)
    }
    
    dfs(root, 0)
    return result
    
}