/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Random *Node
 * }
 */

func copyRandomBinaryTree(root *Node) *NodeCopy {
    origToCp := map[*Node]*NodeCopy{}
    var dfs func(r *Node) *NodeCopy
    dfs = func(r *Node) *NodeCopy {
        // base
        if r == nil {return nil}
        
        // logic
        rootCopy := &NodeCopy{Val: r.Val}
        origToCp[r] = rootCopy
        rootCopy.Left = dfs(r.Left)
        rootCopy.Right = dfs(r.Right)
        return rootCopy
    }
    cp := dfs(root)
    
    var dfsRandomConnector func(r *Node)
    dfsRandomConnector = func(r *Node) {
        // base
        if r == nil {return}
        
        // logic
        if r.Random != nil {
            origToCp[r].Random = origToCp[r.Random]
        }
        dfsRandomConnector(r.Left)
        dfsRandomConnector(r.Right)
    }
    
    dfsRandomConnector(root)
    return cp
}