/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
    level order is also possible
    - level order can be done using bfs or dfs
*/
func maxDepth(root *TreeNode) int {
    if root == nil {return 0}
    q := []*TreeNode{root}
    level := 0
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            if dq.Left != nil {q = append(q, dq.Left)}
            if dq.Right != nil {q = append(q, dq.Right)}
            qSize--
        }
        level++
    }
    return level
}

/*
    bottom-up dfs
    - child will count itself and return back final count to parent
    - parent has to decide which child count take ?
        - left count vs right count?
        - we want max depth, therefore max(leftCount, rightCount)
        - BUT dont forget this parent is a child of some other parent
        - therefore max(leftCount,rightCount)+1
    - when we reach a nil node, a nil child, since there is no node, we return 0
*/
// func maxDepth(root *TreeNode) int {
//     var dfs func(r *TreeNode) int
//     dfs = func(r *TreeNode) int {
//         // base
//         if r == nil {return 0}

//         // logic
//         left := dfs(r.Left)
//         right := dfs(r.Right)
//         return max(left,right)+1
//     }
//     return dfs(root)
// }

/*
    top-down recursion
    - we can keep track of number of nodes in recursion at each node
    - always save the number of nodes when number of nodes in recursion > max seen so far
    time = o(n)
    space = o(n) - skewed tree
*/
// func maxDepth(root *TreeNode) int {
//     if root == nil {return 0}
//     maxN := 0    
//     var dfs func(r *TreeNode, count int) 
//     dfs = func(r *TreeNode, count int) {
//         // base
//         if r == nil {
//             return 
//         }
//         // logic
//         maxN = max(maxN, count)
//         dfs(r.Left, count+1)
//         dfs(r.Right, count+1)
//     }
//     dfs(root, 1)
//     return maxN
// }