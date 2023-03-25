/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Random *Node
 * }
 */
// approach: 1 pass dfs, create random node if orig node has a random ptr and index the random node
// ensure that if random ptr exists, check if we have already created that random node first, use it if have created it already
// otherwise create a new one, index it
// connect it
func copyRandomBinaryTree(root *Node) *NodeCopy {
    origToCp := map[*Node]*NodeCopy{}
    var dfs func(r *Node) *NodeCopy
    dfs = func(r *Node) *NodeCopy {
        // base
        if r == nil {return nil}
        
        // logic
        rootCopy, exists := origToCp[r]
        if !exists {
            rootCopy = &NodeCopy{Val: r.Val}
            origToCp[r] = rootCopy
        }
        if r.Random != nil {
            randomCp, ok := origToCp[r.Random]
            if !ok {
                randomCp = &NodeCopy{Val: r.Random.Val}
                origToCp[r.Random] = randomCp
            }
            rootCopy.Random = randomCp
        }                    
        rootCopy.Left = dfs(r.Left)
        rootCopy.Right = dfs(r.Right)
        return rootCopy
    }
    cp := dfs(root)
    
    return cp
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Random *Node
 * }
 */

// approach: use hashmap to map original node to copy node ( same as copy linked list with random ptr )
// 2 pass ; 1st to create copy tree
// 2nd to connect random ptrs
// func copyRandomBinaryTree(root *Node) *NodeCopy {
//     origToCp := map[*Node]*NodeCopy{}
//     var dfs func(r *Node) *NodeCopy
//     dfs = func(r *Node) *NodeCopy {
//         // base
//         if r == nil {return nil}
        
//         // logic
//         rootCopy := &NodeCopy{Val: r.Val}
//         origToCp[r] = rootCopy
//         rootCopy.Left = dfs(r.Left)
//         rootCopy.Right = dfs(r.Right)
//         return rootCopy
//     }
//     cp := dfs(root)
    
//     var dfsRandomConnector func(r *Node)
//     dfsRandomConnector = func(r *Node) {
//         // base
//         if r == nil {return}
        
//         // logic
//         if r.Random != nil {
//             origToCp[r].Random = origToCp[r.Random]
//         }
//         dfsRandomConnector(r.Left)
//         dfsRandomConnector(r.Right)
//     }
    
//     dfsRandomConnector(root)
//     return cp
// }