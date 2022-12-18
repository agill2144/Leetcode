/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// DFS
// func isCousins(root *TreeNode, x int, y int) bool {
//     var xParent *TreeNode
//     xDepth := -1
//     var yParent *TreeNode
//     yDepth := -1
//     var dfs func(r *TreeNode, parent *TreeNode, depth int) 
//     dfs = func(r *TreeNode, parent *TreeNode, depth int) {
//         // base
//         if r == nil {return}
        
//         // logic
//         if xParent != nil && yParent != nil {return}
        
//         if r.Val == x {
//             xParent = parent
//             xDepth = depth
//         }
//         if r.Val == y {
//             yParent = parent
//             yDepth = depth
//         }
        
//         dfs(r.Left, r, depth+1)
//         dfs(r.Right, r, depth+1)
//     }
    
//     dfs(root, nil, 0)
//     return (xDepth != -1 && yDepth != -1) && (xDepth == yDepth && xParent != yParent)
    
// }


// BFS approach
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
    
    type pair struct {
        node *TreeNode
        parent *TreeNode
    }
    q := []*pair{}
    q = append(q, &pair{node: root, parent: nil})
    var xParent *TreeNode
    var yParent *TreeNode
    
    for len(q) != 0 {
        qSize := len(q)
        for qSize != 0 {
            dq := q[0]
            q = q[1:]
            dqNode := dq.node
            dqParent := dq.parent
            
            if dqNode.Val == x {xParent = dqParent}
            if dqNode.Val == y {yParent = dqParent}
            
            if dqNode.Left != nil {
                q = append(q, &pair{node: dqNode.Left, parent: dqNode})
            }
            if dqNode.Right != nil {
                q = append(q, &pair{node: dqNode.Right, parent: dqNode})
            }
            qSize--
        }
        
        if xParent != nil && yParent != nil {
            return xParent != yParent
        }
        if (xParent != nil && yParent == nil) || (xParent == nil && yParent != nil) {
            return false
        } 
    }
    
    return false
}